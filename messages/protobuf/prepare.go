// Copyright (c) 2019 NEC Laboratories Europe GmbH.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protobuf

import (
	"github.com/golang/protobuf/proto"

	"github.com/hyperledger-labs/minbft/messages"
	"github.com/hyperledger-labs/minbft/messages/protobuf/pb"
)

type prepare struct {
	pbMsg *pb.Prepare
}

func newPrepare(r uint32, v uint64, req messages.Request) *prepare {
	return &prepare{pbMsg: &pb.Prepare{Msg: &pb.Prepare_M{
		ReplicaId: r,
		View:      v,
		Request:   pb.RequestFromAPI(req),
	}}}
}

func newPrepareFromPb(pbMsg *pb.Prepare) *prepare {
	return &prepare{pbMsg: pbMsg}
}

func (m *prepare) MarshalBinary() ([]byte, error) {
	return proto.Marshal(&pb.Message{Type: &pb.Message_Prepare{Prepare: m.pbMsg}})
}

func (m *prepare) ReplicaID() uint32 {
	return m.pbMsg.GetMsg().GetReplicaId()
}

func (m *prepare) View() uint64 {
	return m.pbMsg.GetMsg().GetView()
}

func (m *prepare) Request() messages.Request {
	return newRequestFromPb(m.pbMsg.GetMsg().GetRequest())
}

func (m *prepare) CertifiedPayload() []byte {
	return pb.MarshalOrPanic(m.pbMsg.GetMsg())
}

func (m *prepare) UIBytes() []byte {
	return m.pbMsg.ReplicaUi
}

func (m *prepare) SetUIBytes(uiBytes []byte) {
	m.pbMsg.ReplicaUi = uiBytes
}

func (prepare) ImplementsReplicaMessage() {}
func (prepare) ImplementsPrepare()        {}
