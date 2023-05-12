package api

// SRTStatistics represents the statistics of a SRT connection
type SRTStatistics struct {
	MsTimeStamp uint64 `json:"timestamp_ms" format:"uint64"` // The time elapsed, in milliseconds, since the SRT socket has been created

	// Accumulated

	PktSent         uint64 `json:"sent_pkt" format:"uint64"`           // The total number of sent DATA packets, including retransmitted packets
	PktRecv         uint64 `json:"recv_pkt" format:"uint64"`           // The total number of received DATA packets, including retransmitted packets
	PktSentUnique   uint64 `json:"sent_unique_pkt" format:"uint64"`    // The total number of unique DATA packets sent by the SRT sender
	PktRecvUnique   uint64 `json:"recv_unique_pkt" format:"uint64"`    // The total number of unique original, retransmitted or recovered by the packet filter DATA packets received in time, decrypted without errors and, as a result, scheduled for delivery to the upstream application by the SRT receiver.
	PktSndLoss      uint64 `json:"send_loss_pkt" format:"uint64"`      // The total number of data packets considered or reported as lost at the sender side. Does not correspond to the packets detected as lost at the receiver side.
	PktRcvLoss      uint64 `json:"recv_loss_pkt" format:"uint64"`      // The total number of SRT DATA packets detected as presently missing (either reordered or lost) at the receiver side
	PktRetrans      uint64 `json:"sent_retrans_pkt" format:"uint64"`   // The total number of retransmitted packets sent by the SRT sender
	PktRcvRetrans   uint64 `json:"recv_retran_pkts" format:"uint64"`   // The total number of retransmitted packets registered at the receiver side
	PktSentACK      uint64 `json:"sent_ack_pkt" format:"uint64"`       // The total number of sent ACK (Acknowledgement) control packets
	PktRecvACK      uint64 `json:"recv_ack_pkt" format:"uint64"`       // The total number of received ACK (Acknowledgement) control packets
	PktSentNAK      uint64 `json:"sent_nak_pkt" format:"uint64"`       // The total number of sent NAK (Negative Acknowledgement) control packets
	PktRecvNAK      uint64 `json:"recv_nak_pkt" format:"uint64"`       // The total number of received NAK (Negative Acknowledgement) control packets
	PktSentKM       uint64 `json:"send_km_pkt" format:"uint64"`        // The total number of sent KM (Key Material) control packets
	PktRecvKM       uint64 `json:"recv_km_pkt" format:"uint64"`        // The total number of received KM (Key Material) control packets
	UsSndDuration   uint64 `json:"send_duration_us" format:"uint64"`   // The total accumulated time in microseconds, during which the SRT sender has some data to transmit, including packets that have been sent, but not yet acknowledged
	PktSndDrop      uint64 `json:"send_drop_pkt" format:"uint64"`      // The total number of dropped by the SRT sender DATA packets that have no chance to be delivered in time
	PktRcvDrop      uint64 `json:"recv_drop_pkt" format:"uint64"`      // The total number of dropped by the SRT receiver and, as a result, not delivered to the upstream application DATA packets
	PktRcvUndecrypt uint64 `json:"recv_undecrypt_pkt" format:"uint64"` // The total number of packets that failed to be decrypted at the receiver side

	ByteSent         uint64 `json:"sent_bytes" format:"uint64"`           // Same as pktSent, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecv         uint64 `json:"recv_bytes" format:"uint64"`           // Same as pktRecv, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteSentUnique   uint64 `json:"sent_unique_bytes" format:"uint64"`    // Same as pktSentUnique, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvUnique   uint64 `json:"recv_unique_bytes" format:"uint64"`    // Same as pktRecvUnique, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRcvLoss      uint64 `json:"recv_loss_bytes" format:"uint64"`      // Same as pktRcvLoss, but expressed in bytes, including payload and all the headers (IP, TCP, SRT), bytes for the presently missing (either reordered or lost) packets' payloads are estimated based on the average packet size
	ByteRetrans      uint64 `json:"sent_retrans_bytes" format:"uint64"`   // Same as pktRetrans, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteSndDrop      uint64 `json:"send_drop_bytes" format:"uint64"`      // Same as pktSndDrop, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRcvDrop      uint64 `json:"recv_drop_bytes" format:"uint64"`      // Same as pktRcvDrop, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRcvUndecrypt uint64 `json:"recv_undecrypt_bytes" format:"uint64"` // Same as pktRcvUndecrypt, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)

	// Instantaneous

	UsPktSndPeriod       float64 `json:"pkt_send_period_us"`                           // Current minimum time interval between which consecutive packets are sent, in microseconds
	PktFlowWindow        uint64  `json:"flow_window_pkt" format:"uint64"`              // The maximum number of packets that can be "in flight"
	PktFlightSize        uint64  `json:"flight_size_pkt" format:"uint64"`              // The number of packets in flight
	MsRTT                float64 `json:"rtt_ms"`                                       // Smoothed round-trip time (SRTT), an exponentially-weighted moving average (EWMA) of an endpoint's RTT samples, in milliseconds
	MbpsBandwidth        float64 `json:"bandwidth_mbit"`                               // Estimated bandwidth of the network link, in Mbps
	ByteAvailSndBuf      uint64  `json:"avail_send_buf_bytes" format:"uint64"`         // The available space in the sender's buffer, in bytes
	ByteAvailRcvBuf      uint64  `json:"avail_recv_buf_bytes" format:"uint64"`         // The available space in the receiver's buffer, in bytes
	MbpsMaxBW            float64 `json:"max_bandwidth_mbit"`                           // Transmission bandwidth limit, in Mbps
	ByteMSS              uint64  `json:"mss_bytes" format:"uint64"`                    // Maximum Segment Size (MSS), in bytes
	PktSndBuf            uint64  `json:"send_buf_pkt" format:"uint64"`                 // The number of packets in the sender's buffer that are already scheduled for sending or even possibly sent, but not yet acknowledged
	ByteSndBuf           uint64  `json:"send_buf_bytes" format:"uint64"`               // Instantaneous (current) value of pktSndBuf, but expressed in bytes, including payload and all headers (IP, TCP, SRT)
	MsSndBuf             uint64  `json:"send_buf_ms" format:"uint64"`                  // The timespan (msec) of packets in the sender's buffer (unacknowledged packets)
	MsSndTsbPdDelay      uint64  `json:"send_tsbpd_delay_ms" format:"uint64"`          // Timestamp-based Packet Delivery Delay value of the peer
	PktRcvBuf            uint64  `json:"recv_buf_pkt" format:"uint64"`                 // The number of acknowledged packets in receiver's buffer
	ByteRcvBuf           uint64  `json:"recv_buf_bytes" format:"uint64"`               // Instantaneous (current) value of pktRcvBuf, expressed in bytes, including payload and all headers (IP, TCP, SRT)
	MsRcvBuf             uint64  `json:"recv_buf_ms" format:"uint64"`                  // The timespan (msec) of acknowledged packets in the receiver's buffer
	MsRcvTsbPdDelay      uint64  `json:"recv_tsbpd_delay_ms" format:"uint64"`          // Timestamp-based Packet Delivery Delay value set on the socket via SRTO_RCVLATENCY or SRTO_LATENCY
	PktReorderTolerance  uint64  `json:"reorder_tolerance_pkt" format:"uint64"`        // Instant value of the packet reorder tolerance
	PktRcvAvgBelatedTime uint64  `json:"pkt_recv_avg_belated_time_ms" format:"uint64"` // Accumulated difference between the current time and the time-to-play of a packet that is received late
}

type SRTLog struct {
	Timestamp int64    `json:"ts" format:"int64"`
	Message   []string `json:"msg"`
}

// SRTConnection represents a SRT connection with statistics and logs
type SRTConnection struct {
	Log   map[string][]SRTLog `json:"log"`
	Stats SRTStatistics       `json:"stats"`
}

// SRTChannel represents a SRT publishing connection with its subscribers
type SRTChannel struct {
	Name        string                   `json:"name"`
	SocketId    uint32                   `json:"socketid"`
	Subscriber  []uint32                 `json:"subscriber"`
	Connections map[uint32]SRTConnection `json:"connections"`
	Log         map[string][]SRTLog      `json:"log"`
}
