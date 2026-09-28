package mincostflow

import "testing"

func TestSingleEdge(t *testing.T) {
	flow, cost, err := Solve(2, []Edge{{From: 0, To: 1, Cap: 3, Cost: 2}}, 0, 1)
	if err != nil || flow != 3 || cost != 6 {
		t.Fatalf("单条边不对：%d %d %v", flow, cost, err)
	}
}

func TestNoPath(t *testing.T) {
	flow, cost, err := Solve(3, []Edge{{From: 0, To: 1, Cap: 5, Cost: 1}}, 0, 2)
	if err != nil || flow != 0 || cost != 0 {
		t.Fatalf("没有通路应该给 0：%d %d %v", flow, cost, err)
	}
}

func TestBadInput(t *testing.T) {
	if _, _, err := Solve(0, nil, 0, 0); err != ErrInput {
		t.Fatalf("非法输入应该报 ErrInput：%v", err)
	}
	if _, _, err := Solve(2, []Edge{{From: 0, To: 5, Cap: 1, Cost: 1}}, 0, 1); err != ErrInput {
		t.Fatalf("越界端点应该报 ErrInput：%v", err)
	}
}

func TestSelfLoopIgnored(t *testing.T) {
	edges := []Edge{{From: 0, To: 0, Cap: 9, Cost: 1}, {From: 0, To: 1, Cap: 1, Cost: 1}}
	flow, _, err := Solve(2, edges, 0, 1)
	if err != nil || flow != 1 {
		t.Fatalf("自环应该被忽略：%d %v", flow, err)
	}
}

func TestFlowRespectsCapacity(t *testing.T) {
	edges := []Edge{{From: 0, To: 1, Cap: 2, Cost: 1}, {From: 1, To: 2, Cap: 2, Cost: 1}}
	flow, cost, err := Solve(3, edges, 0, 2)
	if err != nil || flow != 2 || cost != 4 {
		t.Fatalf("直链应该给 2/4：%d %d %v", flow, cost, err)
	}
}
