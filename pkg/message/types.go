package message

import (
	"github.com/sbezverk/gobmp/pkg/sr"
	"github.com/sbezverk/gobmp/pkg/srv6"
)

// PeerStateChange defines a message format sent to as a result of BMP Peer Up or Peer Down message
type PeerStateChange struct {
	Action           string `json:"action"` // Action can be "up" or "down"
	Sequence         int    `json:"sequence,omitempty"`
	Hash             string `json:"hash,omitempty"`
	RouterHash       string `json:"router_hash,omitempty"`
	Name             string `json:"name,omitempty"`
	RemoteBGPID      string `json:"remote_bgp_id,omitempty"`
	RouterIP         string `json:"router_ip,omitempty"`
	Timestamp        string `json:"timestamp,omitempty"`
	RemoteASN        int32  `json:"remote_asn,omitempty"`
	RemoteIP         string `json:"remote_ip,omitempty"`
	PeerRD           string `json:"peer_rd,omitempty"`
	RemotePort       int    `json:"remote_port,omitempty"`
	LocalASN         int32  `json:"local_asn,omitempty"`
	LocalIP          string `json:"local_ip,omitempty"`
	LocalPort        int    `json:"local_port,omitempty"`
	LocalBGPID       string `json:"local_bgp_id,omitempty"`
	InfoData         string `json:"info_data,omitempty"`
	AdvCapabilities  string `json:"adv_cap,omitempty"`
	RcvCapabilities  string `json:"recv_cap,omitempty"`
	RemoteHolddown   int    `json:"remote_holddown,omitempty"`
	AdvHolddown      int    `json:"adv_holddown,omitempty"`
	BMPReason        int    `json:"bmp_reason,omitempty"`
	BMPErrorCode     int    `json:"bmp_error_code,omitempty"`
	BMPErrorSubCode  int    `json:"bmp_error_sub_code,omitempty"`
	ErrorText        string `json:"error_text,omitempty"`
	IsL3VPN          bool   `json:"is_l,omitempty"`
	IsPrepolicy      bool   `json:"isprepolicy,omitempty"`
	IsIPv4           bool   `json:"is_ipv4,omitempty"`
	IsLocRIB         bool   `json:"is_locrib,omitempty"`
	IsLocRIBFiltered bool   `json:"is_locrib_filtered,omitempty"`
	TableName        string `json:"table_name,omitempty"`
}

// UnicastPrefix defines a message format sent as a result of BMP Route Monitor message
// which carries BGP Update with original NLRI information.
type UnicastPrefix struct {
	Action           string   `json:"action"` // Action can be "add" or "del"
	Sequence         int      `json:"sequence,omitempty"`
	Hash             string   `json:"hash,omitempty"`
	RouterHash       string   `json:"router_hash,omitempty"`
	RouterIP         string   `json:"router_ip,omitempty"`
	BaseAttrHash     string   `json:"base_attr_hash,omitempty"`
	PeerHash         string   `json:"peer_hash,omitempty"`
	PeerIP           string   `json:"peer_ip,omitempty"`
	PeerASN          int32    `json:"peer_asn,omitempty"`
	Timestamp        string   `json:"timestamp,omitempty"`
	Prefix           string   `json:"prefix,omitempty"`
	PrefixLen        int32    `json:"prefix_len,omitempty"`
	IsIPv4           bool     `json:"is_ipv4,omitempty"`
	Origin           string   `json:"origin,omitempty"`
	ASPath           []uint32 `json:"as_path,omitempty"`
	ASPathCount      int32    `json:"as_path_count,omitempty"`
	OriginAS         string   `json:"origin_as,omitempty"`
	Nexthop          string   `json:"nexthop,omitempty"`
	MED              uint32   `json:"med,omitempty"`
	LocalPref        uint32   `json:"local_pref,omitempty"`
	Aggregator       string   `json:"aggregator,omitempty"`
	CommunityList    string   `json:"community_list,omitempty"`
	ExtCommunityList string   `json:"ext_community_list,omitempty"`
	IsAtomicAgg      bool     `json:"is_atomic_agg,omitempty"`
	IsNexthopIPv4    bool     `json:"is_nexthop_ipv4,omitempty"`
	OriginatorID     string   `json:"originator_id,omitempty"`
	PathID           int32    `json:"path_id,omitempty"`
	Labels           string   `json:"labels,omitempty"`
	IsPrepolicy      bool     `json:"isprepolicy,omitempty"`
	IsAdjRIBIn       bool     `json:"is_adj_rib_in,omitempty"`
}

// LSNode defines a structure of LS Node message
type LSNode struct {
	Action              string   `json:"action"` // Action can be "add" or "del"
	Sequence            int      `json:"sequence,omitempty"`
	Hash                string   `json:"hash,omitempty"`
	RouterHash          string   `json:"router_hash,omitempty"`
	RouterIP            string   `json:"router_ip,omitempty"`
	BaseAttrHash        string   `json:"base_attr_hash,omitempty"`
	PeerHash            string   `json:"peer_hash,omitempty"`
	PeerIP              string   `json:"peer_ip,omitempty"`
	PeerASN             int32    `json:"peer_asn,omitempty"`
	Timestamp           string   `json:"timestamp,omitempty"`
	IGPRouterID         string   `json:"igp_router_id,omitempty"`
	RouterID            string   `json:"router_id,omitempty"`
	RoutingID           string   `json:"routing_id,omitempty"`
	ASN                 uint32   `json:"asn,omitempty"`
	LSID                uint32   `json:"ls_id,omitempty"`
	MTID                []uint16 `json:"mt_id,omitempty"`
	OSPFAreaID          string   `json:"ospf_area_id,omitempty"`
	ISISAreaID          string   `json:"isis_area_id,omitempty"`
	Protocol            string   `json:"protocol,omitempty"`
	Flags               uint8    `json:"flags,omitempty"`
	ASPath              []uint32 `json:"as_path,omitempty"`
	Nexthop             string   `json:"nexthop,omitempty"`
	MED                 uint32   `json:"med,omitempty"`
	LocalPref           uint32   `json:"local_pref,omitempty"`
	Name                string   `json:"name,omitempty"`
	SRCapabilities      string   `json:"ls_sr_capabilities,omitempty"`
	SRAlgorithm         []int    `json:"sr_algorithm,omitempty"`
	SRLocalBlock        string   `json:"sr_local_block,omitempty"`
	SRv6CapabilitiesTLV string   `json:"srv6_capabilities_tlv,omitempty"`
	NodeMSD             string   `json:"node_msd,omitempty"`
	IsPrepolicy         bool     `json:"isprepolicy,omitempty"`
	IsAdjRIBIn          bool     `json:"is_adj_rib_in,omitempty"`
}

// LSLink defines a structure of LS link message
type LSLink struct {
	Action                string               `json:"action"`
	Sequence              int                  `json:"sequence,omitempty"`
	Hash                  string               `json:"hash,omitempty"`
	RouterHash            string               `json:"router_hash,omitempty"`
	RouterIP              string               `json:"router_ip,omitempty"`
	BaseAttrHash          string               `json:"base_attr_hash,omitempty"`
	PeerHash              string               `json:"peer_hash,omitempty"`
	PeerIP                string               `json:"peer_ip,omitempty"`
	PeerASN               int32                `json:"peer_asn,omitempty"`
	Timestamp             string               `json:"timestamp,omitempty"`
	IGPRouterID           string               `json:"igp_router_id,omitempty"`
	RouterID              string               `json:"router_id,omitempty"`
	RoutingID             string               `json:"routing_id,omitempty"`
	LSID                  uint32               `json:"ls_id,omitempty"`
	OSPFAreaID            string               `json:"ospf_area_id,omitempty"`
	ISISAreaID            string               `json:"isis_area_id,omitempty"`
	Protocol              string               `json:"protocol,omitempty"`
	ASPath                []uint32             `json:"as_path,omitempty"`
	LocalPref             uint32               `json:"local_pref,omitempty"`
	MED                   uint32               `json:"med,omitempty"`
	Nexthop               string               `json:"nexthop,omitempty"`
	MTID                  []uint16             `json:"mt_id,omitempty"`
	LocalLinkID           string               `json:"local_link_id,omitempty"`
	RemoteLinkID          string               `json:"remote_link_id,omitempty"`
	InterfaceIP           string               `json:"intf_ip,omitempty"`
	NeighborIP            string               `json:"nei_ip,omitempty"`
	IGPMetric             uint32               `json:"igp_metric,omitempty"`
	AdminGroup            uint32               `json:"admin_group,omitempty"`
	MaxLinkBW             uint32               `json:"max_link_bw,omitempty"`
	MaxResvBW             uint32               `json:"max_resv_bw,omitempty"`
	UnResvBW              []uint32             `json:"unresv_bw,omitempty"`
	TEDefaultMetric       uint32               `json:"te_default_metric,omitempty"`
	LinkProtection        uint16               `json:"link_protection,omitempty"`
	MPLSProtoMask         uint8                `json:"mpls_proto_mask,omitempty"`
	SRLG                  []uint32             `json:"srlg,omitempty"`
	LinkName              string               `json:"link_name,omitempty"`
	RemoteNodeHash        string               `json:"remote_node_hash,omitempty"`
	LocalNodeHash         string               `json:"local_node_hash,omitempty"`
	RemoteIGPRouterID     string               `json:"remote_igp_router_id,omitempty"`
	RemoteRouterID        string               `json:"remote_router_id,omitempty"`
	LocalNodeASN          uint32               `json:"local_node_asn,omitempty"`
	RemoteNodeASN         uint32               `json:"remote_node_asn,omitempty"`
	SRv6BGPPeerNodeSID    *srv6.BGPPeerNodeSID `json:"srv6_bgp_peer_node_sid,omitempty"`
	IsPrepolicy           bool                 `json:"isprepolicy,omitempty"`
	IsAdjRIBIn            bool                 `json:"is_adj_rib_in,omitempty"`
	LSAdjacencySID        *sr.AdjacencySIDTLV  `json:"ls_adjacency_sid,omitempty"`
	LinkMSD               string               `json:"link_msd,omitempty"`
	UnidirLinkDelay       uint32               `json:"unidir_link_delay,omitempty"`
	UnidirLinkDelayMinMax []uint32             `json:"unidir_link_delay_min_max,omitempty"`
	UnidirDelayVariation  uint32               `json:"unidir_delay_variation,omitempty"`
	UnidirPacketLoss      uint32               `json:"unidir_packet_loss,omitempty"`
	UnidirResidualBW      uint32               `json:"unidir_residual_bw,omitempty"`
	UnidirAvailableBW     uint32               `json:"unidir_available_bw,omitempty"`
	UnidirBWUtilization   uint32               `json:"unidir_bw_utilization,omitempty"`
}

// L3VPNPrefix defines the structure of Layer 3 VPN message
type L3VPNPrefix struct {
	Action           string   `json:"action"` // Action can be "add" or "del"
	Sequence         int      `json:"sequence,omitempty"`
	Hash             string   `json:"hash,omitempty"`
	RouterHash       string   `json:"router_hash,omitempty"`
	RouterIP         string   `json:"router_ip,omitempty"`
	BaseAttrHash     string   `json:"base_attr_hash,omitempty"`
	PeerHash         string   `json:"peer_hash,omitempty"`
	PeerIP           string   `json:"peer_ip,omitempty"`
	PeerASN          int32    `json:"peer_asn,omitempty"`
	Timestamp        string   `json:"timestamp,omitempty"`
	Prefix           string   `json:"prefix,omitempty"`
	PrefixLen        int32    `json:"prefix_len,omitempty"`
	IsIPv4           bool     `json:"is_ipv4,omitempty"`
	Origin           string   `json:"origin,omitempty"`
	ASPath           []uint32 `json:"as_path,omitempty"`
	ASPathCount      int32    `json:"as_path_count,omitempty"`
	OriginAS         string   `json:"origin_as,omitempty"`
	Nexthop          string   `json:"nexthop,omitempty"`
	MED              uint32   `json:"med,omitempty"`
	LocalPref        uint32   `json:"local_pref,omitempty"`
	Aggregator       string   `json:"aggregator,omitempty"`
	CommunityList    string   `json:"community_list,omitempty"`
	ExtCommunityList string   `json:"ext_community_list,omitempty"`
	ClusterList      string   `json:"cluster_list,omitempty"`
	IsAtomicAgg      bool     `json:"is_atomic_agg,omitempty"`
	IsNexthopIPv4    bool     `json:"is_nexthop_ipv4,omitempty"`
	OriginatorID     string   `json:"originator_id,omitempty"`
	PathID           int32    `json:"path_id,omitempty"`
	Labels           []uint32 `json:"labels,omitempty"`
	IsPrepolicy      bool     `json:"isprepolicy,omitempty"`
	IsAdjRIBIn       bool     `json:"is_adj_rib_in,omitempty"`
	VPNRD            string   `json:"vpn_rd,omitempty"`
	VPNRDType        uint16   `json:"vpn_rd_type"`
}

// LSPrefix defines a structure of LS Prefix message
type LSPrefix struct {
	Action        string           `json:"action"`
	Sequence      int              `json:"sequence,omitempty"`
	Hash          string           `json:"hash,omitempty"`
	RouterHash    string           `json:"router_hash,omitempty"`
	RouterIP      string           `json:"router_ip,omitempty"`
	BaseAttrHash  string           `json:"base_attr_hash,omitempty"`
	PeerHash      string           `json:"peer_hash,omitempty"`
	PeerIP        string           `json:"peer_ip,omitempty"`
	PeerASN       int32            `json:"peer_asn,omitempty"`
	Timestamp     string           `json:"timestamp,omitempty"`
	IGPRouterID   string           `json:"igp_router_id,omitempty"`
	RouterID      string           `json:"router_id,omitempty"`
	RoutingID     string           `json:"routing_id,omitempty"`
	LSID          uint32           `json:"ls_id,omitempty"`
	OSPFAreaID    string           `json:"ospf_area_id,omitempty"`
	ISISAreaID    string           `json:"isis_area_id,omitempty"`
	Protocol      string           `json:"protocol,omitempty"`
	ASPath        []uint32         `json:"as_path,omitempty"`
	LocalPref     uint32           `json:"local_pref,omitempty"`
	MED           uint32           `json:"med,omitempty"`
	Nexthop       string           `json:"nexthop,omitempty"`
	LocalNodeHash string           `json:"local_node_hash,omitempty"`
	MTID          []uint16         `json:"mt_id,omitempty"`
	OSPFRouteType uint8            `json:"ospf_route_type,omitempty"`
	IGPFlags      uint8            `json:"igp_flags,omitempty"`
	RouteTag      uint8            `json:"route_tag,omitempty"`
	ExtRouteTag   uint8            `json:"ext_route_tag,omitempty"`
	OSPFFwdAddr   string           `json:"ospf_fwd_addr,omitempty"`
	IGPMetric     uint32           `json:"igp_metric,omitempty"`
	Prefix        string           `json:"prefix,omitempty"`
	PrefixLen     int32            `json:"prefix_len,omitempty"`
	IsPrepolicy   bool             `json:"isprepolicy,omitempty"`
	IsAdjRIBIn    bool             `json:"is_adj_rib_in,omitempty"`
	LSPrefixSID   *sr.PrefixSIDTLV `json:"ls_prefix_sid,omitempty"`
}

// LSSRv6SID defines a structure of LS SRv6 SID message
type LSSRv6SID struct {
	Action               string                 `json:"action"`
	Sequence             int                    `json:"sequence,omitempty"`
	Hash                 string                 `json:"hash,omitempty"`
	RouterHash           string                 `json:"router_hash,omitempty"`
	RouterIP             string                 `json:"router_ip,omitempty"`
	BaseAttrHash         string                 `json:"base_attr_hash,omitempty"`
	PeerHash             string                 `json:"peer_hash,omitempty"`
	PeerIP               string                 `json:"peer_ip,omitempty"`
	PeerASN              int32                  `json:"peer_asn,omitempty"`
	Timestamp            string                 `json:"timestamp,omitempty"`
	IGPRouterID          string                 `json:"igp_router_id,omitempty"`
	LocalNodeASN         uint32                 `json:"local_node_asn,omitempty"`
	RouterID             string                 `json:"router_id,omitempty"`
	RoutingID            string                 `json:"routing_id,omitempty"`
	LSID                 uint32                 `json:"ls_id,omitempty"`
	OSPFAreaID           string                 `json:"ospf_area_id,omitempty"`
	ISISAreaID           string                 `json:"isis_area_id,omitempty"`
	Protocol             string                 `json:"protocol,omitempty"`
	ASPath               []uint32               `json:"as_path,omitempty"`
	LocalPref            uint32                 `json:"local_pref,omitempty"`
	MED                  uint32                 `json:"med,omitempty"`
	Nexthop              string                 `json:"nexthop,omitempty"`
	LocalNodeHash        string                 `json:"local_node_hash,omitempty"`
	MTID                 []uint16               `json:"mt_id,omitempty"`
	OSPFRouteType        uint8                  `json:"ospf_route_type,omitempty"`
	IGPFlags             uint8                  `json:"igp_flags,omitempty"`
	RouteTag             uint8                  `json:"route_tag,omitempty"`
	ExtRouteTag          uint8                  `json:"ext_route_tag,omitempty"`
	OSPFFwdAddr          string                 `json:"ospf_fwd_addr,omitempty"`
	IGPMetric            uint32                 `json:"igp_metric,omitempty"`
	Prefix               string                 `json:"prefix,omitempty"`
	PrefixLen            int32                  `json:"prefix_len,omitempty"`
	IsPrepolicy          bool                   `json:"isprepolicy,omitempty"`
	IsAdjRIBIn           bool                   `json:"is_adj_rib_in,omitempty"`
	SRv6SID              []string               `json:"srv6_sid,omitempty"`
	SRv6EndpointBehavior *srv6.EndpointBehavior `json:"srv6_endpoint_behavior,omitempty"`
	SRv6BGPPeerNodeSID   *srv6.BGPPeerNodeSID   `json:"srv6_bgp_peer_node_sid,omitempty"`
	SRv6SIDStructure     *srv6.SIDStructure     `json:"srv6_sid_structure,omitempty"`
}

// EVPNPrefix defines the structure of EVPN message
type EVPNPrefix struct {
	Action           string   `json:"action"` // Action can be "add" or "del"
	Sequence         int      `json:"sequence,omitempty"`
	Hash             string   `json:"hash,omitempty"`
	RouterHash       string   `json:"router_hash,omitempty"`
	RouterIP         string   `json:"router_ip,omitempty"`
	BaseAttrHash     string   `json:"base_attr_hash,omitempty"`
	PeerHash         string   `json:"peer_hash,omitempty"`
	PeerIP           string   `json:"peer_ip,omitempty"`
	PeerASN          int32    `json:"peer_asn,omitempty"`
	Timestamp        string   `json:"timestamp,omitempty"`
	Prefix           string   `json:"prefix,omitempty"`
	PrefixLen        int32    `json:"prefix_len,omitempty"`
	IsIPv4           bool     `json:"is_ipv4,omitempty"`
	Origin           string   `json:"origin,omitempty"`
	ASPath           []uint32 `json:"as_path,omitempty"`
	ASPathCount      int32    `json:"as_path_count,omitempty"`
	OriginAS         string   `json:"origin_as,omitempty"`
	Nexthop          string   `json:"nexthop,omitempty"`
	MED              uint32   `json:"med,omitempty"`
	LocalPref        uint32   `json:"local_pref,omitempty"`
	Aggregator       string   `json:"aggregator,omitempty"`
	CommunityList    string   `json:"community_list,omitempty"`
	ExtCommunityList string   `json:"ext_community_list,omitempty"`
	ClusterList      string   `json:"cluster_list,omitempty"`
	IsAtomicAgg      bool     `json:"is_atomic_agg,omitempty"`
	IsNexthopIPv4    bool     `json:"is_nexthop_ipv4,omitempty"`
	OriginatorID     string   `json:"originator_id,omitempty"`
	PathID           int32    `json:"path_id,omitempty"`
	Labels           []uint32 `json:"labels,omitempty"`
	IsPrepolicy      bool     `json:"isprepolicy,omitempty"`
	IsAdjRIBIn       bool     `json:"is_adj_rib_in,omitempty"`
	VPNRD            string   `json:"vpn_rd,omitempty"`
	VPNRDType        uint16   `json:"vpn_rd_type"`
}

// 28: (rd):
// 29: (rd_type):
// 30: (rd_type):
// 31: (orig_router_ip_len):
// 32: (eth_tag):
// 33: (eth_segment_id):
// 34: (mac_len):
// 35: (mac):
// 36: (ip_len):
// 37: (ip):
// 38: (label):
// 39: (label):