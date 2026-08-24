package c2q55

import (
	"testing"
)

func TestACL_FailClosedDefault(t *testing.T) {
	mgr := NewACLManager()
	if mgr.CheckAccess("node-a", "telemetry.sensors", PrivilegePublish) {
		t.Fatalf("fail-closed failed: unconfigured ACLManager must deny all")
	}
}

func TestACL_ExactMatchAndWildcard(t *testing.T) {
	mgr := NewACLManager()
	mgr.AddRule("node-a", "telemetry.sensors", PrivilegePublish)
	mgr.AddRule("node-b", "telemetry.*", PrivilegeConsume)
	mgr.AddRule("admin-node", "*", PrivilegeAdmin)

	// node-a : Publish OK, Consume Refused
	if !mgr.CheckAccess("node-a", "telemetry.sensors", PrivilegePublish) {
		t.Fatalf("node-a should have Publish on telemetry.sensors")
	}
	if mgr.CheckAccess("node-a", "telemetry.sensors", PrivilegeConsume) {
		t.Fatalf("node-a should NOT have Consume on telemetry.sensors")
	}
	if mgr.CheckAccess("node-a", "telemetry.logs", PrivilegePublish) {
		t.Fatalf("node-a should NOT have Publish on telemetry.logs")
	}

	// node-b : Consume on telemetry.* OK
	if !mgr.CheckAccess("node-b", "telemetry.sensors", PrivilegeConsume) {
		t.Fatalf("node-b should have Consume on telemetry.sensors")
	}
	if !mgr.CheckAccess("node-b", "telemetry.logs", PrivilegeConsume) {
		t.Fatalf("node-b should have Consume on telemetry.logs")
	}
	if mgr.CheckAccess("node-b", "orders.payment", PrivilegeConsume) {
		t.Fatalf("node-b should NOT have Consume on orders.payment")
	}

	// admin-node : tout autorisé
	if !mgr.CheckAccess("admin-node", "orders.payment", PrivilegePublish|PrivilegeConsume) {
		t.Fatalf("admin-node should have full privileges on any topic")
	}
}

func TestACL_ZeroAlloc(t *testing.T) {
	mgr := NewACLManager()
	mgr.AddRule("node-a", "telemetry.sensors", PrivilegePublish)
	mgr.AddRule("node-b", "telemetry.*", PrivilegeConsume)

	allocs := testing.AllocsPerRun(1000, func() {
		_ = mgr.CheckAccess("node-a", "telemetry.sensors", PrivilegePublish)
	})

	if allocs != 0 {
		t.Fatalf("ACL CheckAccess allocated %f allocs/op, want 0", allocs)
	}
}
