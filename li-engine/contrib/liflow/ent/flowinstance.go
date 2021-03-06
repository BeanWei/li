// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flowdeployment"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flowinstance"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// FlowInstance is the model entity for the FlowInstance schema.
type FlowInstance struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt int64 `json:"-"`
	// FlowDeploymentID holds the value of the "flow_deployment_id" field.
	// 部署的流程ID
	FlowDeploymentID string `json:"flow_deployment_id,omitempty"`
	// RefID holds the value of the "ref_id" field.
	// 引用/调用方ID
	RefID string `json:"ref_id,omitempty"`
	// Status holds the value of the "status" field.
	// 流程状态(1.执行完成 2.执行中 3.执行终止(强制终止))
	Status int8 `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlowInstanceQuery when eager-loading is set.
	Edges FlowInstanceEdges `json:"edges"`
}

// FlowInstanceEdges holds the relations/edges for other nodes in the graph.
type FlowInstanceEdges struct {
	// FlowDeployment holds the value of the flow_deployment edge.
	FlowDeployment *FlowDeployment `json:"flow_deployment,omitempty"`
	// FlowNodeInstances holds the value of the flow_node_instances edge.
	FlowNodeInstances []*FlowNodeInstance `json:"flow_node_instances,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FlowDeploymentOrErr returns the FlowDeployment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlowInstanceEdges) FlowDeploymentOrErr() (*FlowDeployment, error) {
	if e.loadedTypes[0] {
		if e.FlowDeployment == nil {
			// The edge flow_deployment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: flowdeployment.Label}
		}
		return e.FlowDeployment, nil
	}
	return nil, &NotLoadedError{edge: "flow_deployment"}
}

// FlowNodeInstancesOrErr returns the FlowNodeInstances value or an error if the edge
// was not loaded in eager-loading.
func (e FlowInstanceEdges) FlowNodeInstancesOrErr() ([]*FlowNodeInstance, error) {
	if e.loadedTypes[1] {
		return e.FlowNodeInstances, nil
	}
	return nil, &NotLoadedError{edge: "flow_node_instances"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FlowInstance) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case flowinstance.FieldCreatedAt, flowinstance.FieldUpdatedAt, flowinstance.FieldDeletedAt, flowinstance.FieldStatus:
			values[i] = new(sql.NullInt64)
		case flowinstance.FieldID, flowinstance.FieldFlowDeploymentID, flowinstance.FieldRefID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FlowInstance", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FlowInstance fields.
func (fi *FlowInstance) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flowinstance.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				fi.ID = value.String
			}
		case flowinstance.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fi.CreatedAt = value.Int64
			}
		case flowinstance.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fi.UpdatedAt = value.Int64
			}
		case flowinstance.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				fi.DeletedAt = value.Int64
			}
		case flowinstance.FieldFlowDeploymentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flow_deployment_id", values[i])
			} else if value.Valid {
				fi.FlowDeploymentID = value.String
			}
		case flowinstance.FieldRefID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref_id", values[i])
			} else if value.Valid {
				fi.RefID = value.String
			}
		case flowinstance.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				fi.Status = int8(value.Int64)
			}
		}
	}
	return nil
}

// QueryFlowDeployment queries the "flow_deployment" edge of the FlowInstance entity.
func (fi *FlowInstance) QueryFlowDeployment() *FlowDeploymentQuery {
	return (&FlowInstanceClient{config: fi.config}).QueryFlowDeployment(fi)
}

// QueryFlowNodeInstances queries the "flow_node_instances" edge of the FlowInstance entity.
func (fi *FlowInstance) QueryFlowNodeInstances() *FlowNodeInstanceQuery {
	return (&FlowInstanceClient{config: fi.config}).QueryFlowNodeInstances(fi)
}

// Update returns a builder for updating this FlowInstance.
// Note that you need to call FlowInstance.Unwrap() before calling this method if this FlowInstance
// was returned from a transaction, and the transaction was committed or rolled back.
func (fi *FlowInstance) Update() *FlowInstanceUpdateOne {
	return (&FlowInstanceClient{config: fi.config}).UpdateOne(fi)
}

// Unwrap unwraps the FlowInstance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fi *FlowInstance) Unwrap() *FlowInstance {
	tx, ok := fi.config.driver.(*txDriver)
	if !ok {
		panic("ent: FlowInstance is not a transactional entity")
	}
	fi.config.driver = tx.drv
	return fi
}

// String implements the fmt.Stringer.
func (fi *FlowInstance) String() string {
	var builder strings.Builder
	builder.WriteString("FlowInstance(")
	builder.WriteString(fmt.Sprintf("id=%v", fi.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", fi.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", fi.UpdatedAt))
	builder.WriteString(", deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", fi.DeletedAt))
	builder.WriteString(", flow_deployment_id=")
	builder.WriteString(fi.FlowDeploymentID)
	builder.WriteString(", ref_id=")
	builder.WriteString(fi.RefID)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", fi.Status))
	builder.WriteByte(')')
	return builder.String()
}

type (
	IFlowInstance struct {
		ID                string               `json:"id,omitempty"`
		CreatedAt         int64                `json:"created_at,omitempty"`
		UpdatedAt         int64                `json:"updated_at,omitempty"`
		DeletedAt         int64                `json:"-"`
		FlowDeploymentID  string               `json:"flow_deployment_id,omitempty"`
		RefID             string               `json:"ref_id,omitempty"`
		Status            int8                 `json:"status,omitempty"`
		FlowDeployment    *IFlowDeployment     `json:"flow_deployment,omitempty"`
		FlowNodeInstances []*IFlowNodeInstance `json:"flow_node_instances,omitempty"`
	}
	ListFlowInstanceReq struct {
		Page   int                     `json:"page" d:"1" v:"min:1"`
		Limit  int                     `json:"limit" d:"20" v:"min:1|max:500"`
		Query  string                  `json:"query"`
		Filter *ListFlowInstanceFilter `json:"filter"`
		Sorter *ListFlowInstanceSorter `json:"sorter"`
	}
	ListFlowInstanceFilter struct {
	}
	ListFlowInstanceSorter struct {
	}
	ListFlowInstanceRes struct {
		List  []*IFlowInstance `json:"list"`
		Total int              `json:"total"`
	}
	CreateFlowInstanceReq struct {
	}
	GetFlowInstanceReq struct {
		ID string `json:"id" v:"required"`
	}
	UpdateFlowInstanceReq struct {
		ID string `json:"id" v:"required"`
	}
	DeleteFlowInstanceReq struct {
		ID string `json:"id" v:"required"`
	}
	DeleteManyFlowInstanceReq struct {
		IDs []string `json:"ids" v:"required"`
	}
)

func NewIFlowInstance(e *FlowInstance) *IFlowInstance {
	if e == nil {
		return nil
	}
	return &IFlowInstance{
		ID:                e.ID,
		CreatedAt:         e.CreatedAt,
		UpdatedAt:         e.UpdatedAt,
		DeletedAt:         e.DeletedAt,
		FlowDeploymentID:  e.FlowDeploymentID,
		RefID:             e.RefID,
		Status:            e.Status,
		FlowDeployment:    NewIFlowDeployment(e.Edges.FlowDeployment),
		FlowNodeInstances: NewIFlowNodeInstanceArray(e.Edges.FlowNodeInstances),
	}
}

func NewIFlowInstanceArray(es []*FlowInstance) []*IFlowInstance {
	if len(es) == 0 {
		return nil
	}
	r := make([]*IFlowInstance, len(es))
	for i, e := range es {
		r[i] = NewIFlowInstance(e)
	}
	return r
}

func ListFlowInstanceController(ctx context.Context, req *ListFlowInstanceReq) (res *ListFlowInstanceRes, err error) {
	q := DB().FlowInstance.Query()
	if req.Filter != nil {
	}
	res = &ListFlowInstanceRes{}
	res.Total, err = q.Count(ctx)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	ret, err := q.Limit(req.Limit).Offset((req.Page - 1) * req.Limit).All(ctx)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	res.List = NewIFlowInstanceArray(ret)
	return res, nil
}

func CreateFlowInstanceController(ctx context.Context, req *CreateFlowInstanceReq) (err error) {
	b := DB().FlowInstance.Create()
	err = b.Exec(ctx)
	if err != nil {
		if IsConstraintError(err) {
			return gerror.WrapCode(gcode.CodeOperationFailed, err)
		}
		return gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	return
}

func GetFlowInstanceController(ctx context.Context, req *GetFlowInstanceReq) (res *IFlowInstance, err error) {
	ret, err := DB().FlowInstance.
		Query().
		Where(flowinstance.IDEQ(req.ID)).
		Only(ctx)
	if err != nil {
		if IsNotFound(err) {
			return nil, gerror.WrapCode(gcode.CodeNotFound, err)
		}
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	return NewIFlowInstance(ret), nil
}

func UpdateFlowInstanceController(ctx context.Context, req *UpdateFlowInstanceReq) (err error) {
	b := DB().FlowInstance.UpdateOneID(req.ID)
	err = b.Exec(ctx)
	if err != nil {
		if IsNotFound(err) {
			return gerror.WrapCode(gcode.CodeNotFound, err)
		} else if IsConstraintError(err) {
			return gerror.WrapCode(gcode.CodeOperationFailed, err)
		}
		return gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	return
}

func DeleteFlowInstanceController(ctx context.Context, req *DeleteFlowInstanceReq) (err error) {
	err = DB().FlowInstance.DeleteOneID(req.ID).Exec(ctx)
	if err != nil {
		if IsNotFound(err) {
			return gerror.WrapCode(gcode.CodeNotFound, err)
		} else if IsConstraintError(err) {
			return gerror.WrapCode(gcode.CodeOperationFailed, err)
		}
		return gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	return
}

func DeleteManyFlowInstanceController(ctx context.Context, req *DeleteManyFlowInstanceReq) (err error) {
	_, err = DB().FlowInstance.Delete().Where(flowinstance.IDIn(req.IDs...)).Exec(ctx)
	if err != nil {
		return gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	return
}

func CreateFlowInstanceFormView() view.Node {
	return node.FormGrid("grid").MaxColumns(2).Children()
}

func ReadFlowInstanceFormView() view.Node {
	return node.FormGrid("grid").MaxColumns(2).Children()
}

func UpdateFlowInstanceFormView() view.Node {
	return node.FormGrid("grid").MaxColumns(2).Children()
}

func ListFlowInstanceTableColumns() []view.Node {
	return []view.Node{}
}

func ListFlowInstanceView() view.Node {
	return node.List("flowinstanceList").
		AC(flowinstanceViewACL["list:FlowInstance"]).
		ForInit("@listFlowInstance", ListFlowInstanceController).
		DecoratorCard().
		EnableFilter().
		SelectionMultiple(true).
		Children(
			node.ListTable("flowinstanceListTable").
				ActionBar(
					node.ListAction("flowinstanceListActions").Children(
						node.ListActionRecordFormDrawer("addFlowInstance").
							Title("addNew").
							AC(flowinstanceViewACL["create:FlowInstance"]).
							ButtonType("primary").
							ButtonIcon("IconPlus").
							ButtonPosition("left").
							Body(CreateFlowInstanceFormView()).
							Footer(
								node.ActionFormDrawerCancel("cancel"),
								node.ActionFormDrawerSubmit("submit").ForSubmit("@addFlowInstance", CreateFlowInstanceController),
							),
						node.ListActionRowSelection("deleteManyFlowInstance").
							Title("bulkDelete").
							AC(flowinstanceViewACL["deleteMany:FlowInstance"]).
							ForSubmit("@deleteManyFlowInstance", DeleteManyFlowInstanceController).
							AfterReload(true).
							ConfirmTitle("confirmDelete").
							ButtonStatus("danger").
							ButtonIcon("IconDelete").
							ButtonPosition("left"),
					),
				).
				Columns(
					append(
						ListFlowInstanceTableColumns(),
						node.ListTableColumn("columnAction").
							Title("columnAction").
							DataIndex("__action").
							Width(150).
							Render(
								node.Space("actions").Size(0).SplitByDivider().Children(
									node.ListActionRecordFormDrawer("view").
										AC(flowinstanceViewACL["get:FlowInstance"]).
										ForInit("@getFlowInstance", GetFlowInstanceController).
										DrawerTitle("viewDrawerTitle").
										ButtonType("text").
										ButtonIcon("IconEye").
										Body(ReadFlowInstanceFormView()).
										Footer(
											node.ActionFormDrawerCancel("cancel"),
										),
									node.ListActionRecordFormDrawer("edit").
										AC(flowinstanceViewACL["update:FlowInstance"]).
										ForInit("@getFlowInstance", GetFlowInstanceController).
										DrawerTitle("editDrawerTitle").
										ButtonType("text").
										ButtonIcon("IconEdit").
										Body(UpdateFlowInstanceFormView()).
										Footer(
											node.ActionFormDrawerCancel("cancel"),
											node.ActionFormDrawerSubmit("submit").
												ForSubmit("@updateFlowInstance", UpdateFlowInstanceController),
										),
									node.ListActionRecordDelete("delete").
										AC(flowinstanceViewACL["delete:FlowInstance"]).
										ForSubmit("@deleteFlowInstance", DeleteFlowInstanceController).
										ButtonType("text").
										ButtonIcon("IconDelete"),
								),
							),
					)...,
				),
		)
}

var flowinstanceViewACL = map[string]ac.AC{}

func init() {
	rf := reflect.ValueOf(schema.FlowInstance{}).MethodByName("ACL")
	if rf.IsValid() {
		flowinstanceViewACL, _ = rf.Call([]reflect.Value{})[0].Interface().(map[string]ac.AC)
	}
}

// FlowInstances is a parsable slice of FlowInstance.
type FlowInstances []*FlowInstance

func (fi FlowInstances) config(cfg config) {
	for _i := range fi {
		fi[_i].config = cfg
	}
}
