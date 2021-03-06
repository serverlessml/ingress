/*
Copyright © 2020 Dmitry Kisler <admin@dkisler.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// +build !aws

package bus

import (
	"context"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

// GCPClient defines the client to communicate with pubsub.
type GCPClient struct {
	// GCP Project ID
	ProjectID string
	Ctx       context.Context
	Opts      []option.ClientOption
	Instance  *pubsub.Client
}

// Connect establishes connector to the message broker.
func (c *GCPClient) Connect() error {
	var err error
	c.Ctx = context.Background()
	c.Instance, err = pubsub.NewClient(c.Ctx, c.ProjectID, c.Opts...)
	return err
}

// Push pushes the message to a topic.
func (c *GCPClient) Push(payload []byte, topic string) error {
	t := c.Instance.Topic(topic)
	t.PublishSettings.NumGoroutines = 1

	result := t.Publish(c.Ctx, &pubsub.Message{Data: payload})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	_, err := result.Get(c.Ctx)
	return err
}
