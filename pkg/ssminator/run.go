package ssminator

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

// Command is a copy of "github.com/aws/aws-sdk-go-v2/service/ssm/types.Command"
type Command struct {
	CommandID      string
	CompletedCount int32
	ErrorCount     int32
	StatusDetails  string
}

// RunOnAllInstances run a given command on all available fleet-managed instances.
// Returns asynchronously. The status of the job needs to be queried based on CommandID.
// TODOs:
// - allow passing of params
func (s ssminator) RunOnAllInstances() (*Command, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	instanceIDs, err := s.getManagedInstanceIDList()
	if err != nil {
		return nil, err
	}

	docParams := map[string][]string{
		`commands`: {"uname -a"},
	}

	params := &ssm.SendCommandInput{
		// https://ap-southeast-2.console.aws.amazon.com/systems-manager/documents/AWS-RunShellScript/description
		DocumentName:     aws.String("AWS-RunShellScript"), // arg
		DocumentHash:     aws.String("99749de5e62f71e5ebe9a55c2321e2c394796afe7208cff048696541e6f6771e"),
		DocumentHashType: types.DocumentHashTypeSha256,
		DocumentVersion:  aws.String("1"),
		Parameters:       docParams,
		Comment:          aws.String("Test"), // arg
		InstanceIds:      instanceIDs,        // TODO
		MaxConcurrency:   aws.String("100%"), // arg
		MaxErrors:        aws.String("100%"), // arg
		TimeoutSeconds:   aws.Int32(600),     // arg
	}
	out, err := s.ssmClient.SendCommand(ctx, params)
	if err != nil {
		return nil, err
	}

	command := &Command{
		CommandID:      *out.Command.CommandId,
		CompletedCount: out.Command.CompletedCount,
		ErrorCount:     out.Command.ErrorCount,
		StatusDetails:  *out.Command.StatusDetails,
	}

	return command, nil
}

func (s ssminator) getManagedInstanceIDList() ([]string, error) {
	ii, err := s.DescribeInstanceInformation()
	if err != nil {
		return nil, err
	}
	instanceIDList := make([]string, 0)
	for _, i := range ii {
		instanceIDList = append(instanceIDList, i.InstanceID)
	}
	return instanceIDList, nil
}

func (s ssminator) CheckCommandStatus(cmdID string) (*Command, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	params := &ssm.ListCommandsInput{
		CommandId: aws.String(cmdID),
	}

	cmds, err := s.ssmClient.ListCommands(ctx, params)
	if err != nil {
		return nil, err
	}
	c := cmds.Commands[0] // can only be one, since we provided a CommandID in params
	cmdInvOut := &Command{
		CommandID:      *c.CommandId,
		CompletedCount: c.CompletedCount,
		ErrorCount:     c.ErrorCount,
		StatusDetails:  *c.StatusDetails,
	}
	return cmdInvOut, nil
}
