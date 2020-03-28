package message

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/golang/glog"
	"github.com/sbezverk/gobmp/pkg/bmp"
)

func (p *producer) producePeerUpMessage(msg bmp.Message) {
	if msg.PeerHeader == nil {
		glog.Errorf("perPeerHeader is missing, cannot construct PeerStateChange message")
		return
	}
	peerUpMsg, ok := msg.Payload.(*bmp.PeerUpMessage)
	if !ok {
		glog.Errorf("got invalid Payload type in bmp.Message")
		return
	}

	m := PeerStateChange{
		Action:         "up",
		RouterHash:     msg.PeerHeader.GetPeerHash(),
		RemoteASN:      int16(msg.PeerHeader.PeerAS),
		PeerRD:         msg.PeerHeader.PeerDistinguisher.String(),
		RemotePort:     int(peerUpMsg.RemotePort),
		Timestamp:      msg.PeerHeader.PeerTimestamp,
		LocalASN:       int(peerUpMsg.SentOpen.MyAS),
		LocalPort:      int(peerUpMsg.LocalPort),
		AdvHolddown:    int(peerUpMsg.SentOpen.HoldTime),
		RemoteHolddown: int(peerUpMsg.ReceivedOpen.HoldTime),
	}
	if msg.PeerHeader.FlagV {
		m.IsIPv4 = false
		m.RemoteIP = net.IP(msg.PeerHeader.PeerAddress).To16().String()
		m.LocalIP = net.IP(peerUpMsg.LocalAddress).To16().String()
		m.RemoteBGPID = net.IP(msg.PeerHeader.PeerBGPID).To16().String()
		m.LocalBGPID = net.IP(peerUpMsg.SentOpen.BGPID).To16().String()
	} else {
		m.IsIPv4 = true
		m.RemoteIP = net.IP(msg.PeerHeader.PeerAddress[12:]).To4().String()
		m.LocalIP = net.IP(peerUpMsg.LocalAddress[12:]).To4().String()
		m.RemoteBGPID = net.IP(msg.PeerHeader.PeerBGPID).To4().String()
		m.LocalBGPID = net.IP(peerUpMsg.SentOpen.BGPID).To4().String()
	}

	sCaps := peerUpMsg.SentOpen.GetCapabilities()
	rCaps := peerUpMsg.ReceivedOpen.GetCapabilities()
	for i, cap := range sCaps {
		m.AdvCapabilities += cap.Description
		if i < len(sCaps)-1 {
			m.AdvCapabilities += ", "
		}
	}
	for i, cap := range rCaps {
		m.RcvCapabilities += cap.Description
		if i < len(rCaps)-1 {
			m.RcvCapabilities += ", "
		}
	}
	j, err := json.Marshal(&m)
	if err != nil {
		glog.Errorf("failed to Marshal PeerStateChange struct with error: %+v", err)
		return
	}
	if err := p.publisher.PublishMessage(bmp.PeerUpMsg, []byte(m.RouterHash), j); err != nil {
		glog.Errorf("failed to push PeerUp message to kafka with error: %+v", err)
		return
	}

	glog.V(5).Infof("succeeded to push PeerUp message to kafka")
}

func (p *producer) producePeerDownMessage(msg bmp.Message) {
	if msg.PeerHeader == nil {
		glog.Errorf("perPeerHeader is missing, cannot construct PeerStateChange message")
		return
	}
	peerDownMsg, ok := msg.Payload.(*bmp.PeerDownMessage)
	if !ok {
		glog.Errorf("got invalid Payload type in bmp.Message")
		return
	}
	m := PeerStateChange{
		Action:     "down",
		BMPReason:  int(peerDownMsg.Reason),
		RouterHash: msg.PeerHeader.GetPeerHash(),
		RemoteASN:  int16(msg.PeerHeader.PeerAS),
		PeerRD:     msg.PeerHeader.PeerDistinguisher.String(),
		Timestamp:  msg.PeerHeader.PeerTimestamp,
	}
	if msg.PeerHeader.FlagV {
		m.IsIPv4 = false
		m.RemoteIP = net.IP(msg.PeerHeader.PeerAddress).To16().String()
		m.RemoteBGPID = net.IP(msg.PeerHeader.PeerBGPID).To16().String()
	} else {
		m.IsIPv4 = true
		m.RemoteIP = net.IP(msg.PeerHeader.PeerAddress[12:]).To4().String()
		m.RemoteBGPID = net.IP(msg.PeerHeader.PeerBGPID).To4().String()
	}
	m.InfoData = fmt.Sprintf("%s", peerDownMsg.Data)

	j, err := json.Marshal(&m)
	if err != nil {
		glog.Errorf("failed to Marshal PeerStateChange struct with error: %+v", err)
		return
	}
	if err := p.publisher.PublishMessage(bmp.PeerDownMsg, []byte(m.RouterHash), j); err != nil {
		glog.Errorf("failed to push PeerDown message to kafka with error: %+v", err)
		return
	}

	glog.V(5).Infof("succeeded to push PeerDown message to kafka")
}
