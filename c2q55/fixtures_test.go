// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"encoding/json"
	"fmt"
	"time"
)

// SecurityTelemetryEvent représente un événement réel de sonde de cyberdéfense (c2blue55 / eBPF).
type SecurityTelemetryEvent struct {
	TimestampNs uint64 `json:"timestamp_ns"`
	ProcessID   uint32 `json:"pid"`
	UserID      uint32 `json:"uid"`
	SyscallName string `json:"syscall"`
	SourceIP    string `json:"src_ip"`
	DestIP      string `json:"dst_ip"`
	DestPort    uint16 `json:"dst_port"`
	SHA256Hash  string `json:"sha256"`
	Severity    string `json:"severity"`
	AuditAction string `json:"action"`
}

// OrderTransactionEvent représente un événement réel de transaction commerciale / financière.
type OrderTransactionEvent struct {
	OrderID      string  `json:"order_id"`
	CustomerID   string  `json:"customer_id"`
	SKU          string  `json:"sku"`
	Quantity     uint32  `json:"qty"`
	UnitPriceEUR float64 `json:"unit_price_eur"`
	Currency     string  `json:"currency"`
	PaymentState string  `json:"payment_state"`
	TraceID      string  `json:"trace_id"`
}

// MakeRealisticTelemetryPayload génère une charge utile JSON réelle de 256 à 512 octets.
func MakeRealisticTelemetryPayload(idx int) []byte {
	evt := SecurityTelemetryEvent{
		TimestampNs: uint64(time.Now().UnixNano()),
		ProcessID:   uint32(1000 + (idx % 250)),
		UserID:      1001,
		SyscallName: "sys_enter_execve",
		SourceIP:    fmt.Sprintf("192.168.1.%d", 10+(idx%100)),
		DestIP:      "10.0.55.42",
		DestPort:    8155,
		SHA256Hash:  fmt.Sprintf("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852%04d", idx%10000),
		Severity:    "HIGH_PRIORITY",
		AuditAction: "INTERCEPT_AND_JOURNAL",
	}
	bytes, _ := json.Marshal(evt)
	return bytes
}

// MakeRealisticOrderPayload génère un événement de transaction binaire / JSON réel.
func MakeRealisticOrderPayload(idx int) []byte {
	evt := OrderTransactionEvent{
		OrderID:      fmt.Sprintf("01a02fe2-9a1d-7302-b584-%012d", idx),
		CustomerID:   fmt.Sprintf("cust-eu-west-%06d", idx%5000),
		SKU:          "SRV-NVME-PCIE4-64G",
		Quantity:     uint32(1 + (idx % 8)),
		UnitPriceEUR: 1249.50,
		Currency:     "EUR",
		PaymentState: "SETTLED_SEPA_INSTANT",
		TraceID:      fmt.Sprintf("tr-c2q55-node-%08d", idx),
	}
	bytes, _ := json.Marshal(evt)
	return bytes
}

// MakeRealisticDocumentPayload génère un document de 64 Ko à 1 Mo avec du texte riche et des structures d'audit réelles.
func MakeRealisticDocumentPayload(targetSize int) []byte {
	baseEvent := MakeRealisticTelemetryPayload(42)
	buf := make([]byte, 0, targetSize)

	for len(buf) < targetSize {
		idx := len(buf) / len(baseEvent)
		evt := MakeRealisticOrderPayload(idx)
		buf = append(buf, evt...)
		buf = append(buf, '\n')
	}

	if len(buf) > targetSize {
		buf = buf[:targetSize]
	}
	return buf
}
