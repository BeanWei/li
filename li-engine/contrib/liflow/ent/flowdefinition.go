// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flowdefinition"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// FlowDefinition is the model entity for the FlowDefinition schema.
type FlowDefinition struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt int64 `json:"-"`
	// Name holds the value of the "name" field.
	// 流程名称
	Name string `json:"name,omitempty"`
	// Status holds the value of the "status" field.
	// 状态(1.初始态 1.编辑中 2.已上线)
	Status int8 `json:"status,omitempty"`
	// Model holds the value of the "model" field.
	// 流程模型
	Model schema.FlowModel `json:"model,omitempty"`
	// Remark holds the value of the "remark" field.
	// 备注
	Remark string `json:"remark,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlowDefinitionQuery when eager-loading is set.
	Edges FlowDefinitionEdges `json:"edges"`
}

// FlowDefinitionEdges holds the relations/edges for other nodes in the graph.
type FlowDefinitionEdges struct {
	// FlowDeployments holds the value of the flow_deployments edge.
	FlowDeployments []*FlowDeployment `json:"flow_deployments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FlowDeploymentsOrErr returns the FlowDeployments value or an error if the edge
// was not loaded in eager-loading.
func (e FlowDefinitionEdges) FlowDeploymentsOrErr() ([]*FlowDeployment, error) {
	if e.loadedTypes[0] {
		return e.FlowDeployments, nil
	}
	return nil, &NotLoadedError{edge: "flow_deployments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FlowDefinition) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case flowdefinition.FieldModel:
			values[i] = new([]byte)
		case flowdefinition.FieldCreatedAt, flowdefinition.FieldUpdatedAt, flowdefinition.FieldDeletedAt, flowdefinition.FieldStatus:
			values[i] = new(sql.NullInt64)
		case flowdefinition.FieldID, flowdefinition.FieldName, flowdefinition.FieldRemark:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FlowDefinition", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FlowDefinition fields.
func (fd *FlowDefinition) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flowdefinition.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				fd.ID = value.String
			}
		case flowdefinition.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fd.CreatedAt = value.Int64
			}
		case flowdefinition.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fd.UpdatedAt = value.Int64
			}
		case flowdefinition.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				fd.DeletedAt = value.Int64
			}
		case flowdefinition.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				fd.Name = value.String
			}
		case flowdefinition.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				fd.Status = int8(value.Int64)
			}
		case flowdefinition.FieldModel:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &fd.Model); err != nil {
					return fmt.Errorf("unmarshal field model: %w", err)
				}
			}
		case flowdefinition.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				fd.Remark = value.String
			}
		}
	}
	return nil
}

// QueryFlowDeployments queries the "flow_deployments" edge of the FlowDefinition entity.
func (fd *FlowDefinition) QueryFlowDeployments() *FlowDeploymentQuery {
	return (&FlowDefinitionClient{config: fd.config}).QueryFlowDeployments(fd)
}

// Update returns a builder for updating this FlowDefinition.
// Note that you need to call FlowDefinition.Unwrap() before calling this method if this FlowDefinition
// was returned from a transaction, and the transaction was committed or rolled back.
func (fd *FlowDefinition) Update() *FlowDefinitionUpdateOne {
	return (&FlowDefinitionClient{config: fd.config}).UpdateOne(fd)
}

// Unwrap unwraps the FlowDefinition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fd *FlowDefinition) Unwrap() *FlowDefinition {
	tx, ok := fd.config.driver.(*txDriver)
	if !ok {
		panic("ent: FlowDefinition is not a transactional entity")
	}
	fd.config.driver = tx.drv
	return fd
}

// String implements the fmt.Stringer.
func (fd *FlowDefinition) String() string {
	var builder strings.Builder
	builder.WriteString("FlowDefinition(")
	builder.WriteString(fmt.Sprintf("id=%v", fd.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", fd.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", fd.UpdatedAt))
	builder.WriteString(", deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", fd.DeletedAt))
	builder.WriteString(", name=")
	builder.WriteString(fd.Name)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", fd.Status))
	builder.WriteString(", model=")
	builder.WriteString(fmt.Sprintf("%v", fd.Model))
	builder.WriteString(", remark=")
	builder.WriteString(fd.Remark)
	builder.WriteByte(')')
	return builder.String()
}

type (
	IFlowDefinition struct {
		ID              string             `json:"id,omitempty"`
		CreatedAt       int64              `json:"created_at,omitempty"`
		UpdatedAt       int64              `json:"updated_at,omitempty"`
		DeletedAt       int64              `json:"-"`
		Name            string             `json:"name,omitempty"`
		Status          int8               `json:"status,omitempty"`
		Model           schema.FlowModel   `json:"model,omitempty"`
		Remark          string             `json:"remark,omitempty"`
		FlowDeployments []*IFlowDeployment `json:"flow_deployments,omitempty"`
	}
	ListFlowDefinitionReq struct {
		Page   int                       `json:"page" d:"1" v:"min:1"`
		Limit  int                       `json:"limit" d:"20" v:"min:1|max:500"`
		Query  string                    `json:"query"`
		Filter *ListFlowDefinitionFilter `json:"filter"`
		Sorter *ListFlowDefinitionSorter `json:"sorter"`
	}
	ListFlowDefinitionFilter struct {
	}
	ListFlowDefinitionSorter struct {
	}
	ListFlowDefinitionRes struct {
		List  []*IFlowDefinition `json:"list"`
		Total int                `json:"total"`
	}
	CreateFlowDefinitionReq struct {
	}
	GetFlowDefinitionReq struct {
		ID string `json:"id" v:"required"`
	}
	UpdateFlowDefinitionReq struct {
		ID string `json:"id" v:"required"`
	}
	DeleteFlowDefinitionReq struct {
		ID string `json:"id" v:"required"`
	}
	DeleteManyFlowDefinitionReq struct {
		IDs []string `json:"ids" v:"required"`
	}
)

func NewIFlowDefinition(e *FlowDefinition) *IFlowDefinition {
	if e == nil {
		return nil
	}
	return &IFlowDefinition{
		ID:              e.ID,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
		DeletedAt:       e.DeletedAt,
		Name:            e.Name,
		Status:          e.Status,
		Model:           e.Model,
		Remark:          e.Remark,
		FlowDeployments: NewIFlowDeploymentArray(e.Edges.FlowDeployments),
	}
}

func NewIFlowDefinitionArray(es []*FlowDefinition) []*IFlowDefinition {
	if len(es) == 0 {
		return nil
	}
	r := make([]*IFlowDefinition, len(es))
	for i, e := range es {
		r[i] = NewIFlowDefinition(e)
	}
	return r
}

func ListFlowDefinitionController(ctx context.Context, req *ListFlowDefinitionReq) (res *ListFlowDefinitionRes, err error) {
	q := DB().FlowDefinition.Query()
	if req.Filter != nil {
	}
	res = &ListFlowDefinitionRes{}
	res.Total, err = q.Count(ctx)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	ret, err := q.Limit(req.Limit).Offset((req.Page - 1) * req.Limit).All(ctx)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	res.List = NewIFlowDefinitionArray(ret)
	return res, nil
}

func CreateFlowDefinitionController(ctx context.Context, req *CreateFlowDefinitionReq) (err error) {
	b := DB().FlowDefinition.Create()
	err = b.Exec(ctx)
	if err != nil {
		if IsConstraintError(err) {
			return gerror.WrapCode(gcode.CodeOperationFailed, err)
		}
		return gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	return
}

func GetFlowDefinitionController(ctx context.Context, req *GetFlowDefinitionReq) (res *IFlowDefinition, err error) {
	ret, err := DB().FlowDefinition.
		Query().
		Where(flowdefinition.IDEQ(req.ID)).
		Only(ctx)
	if err != nil {
		if IsNotFound(err) {
			return nil, gerror.WrapCode(gcode.CodeNotFound, err)
		}
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	return NewIFlowDefinition(ret), nil
}

func UpdateFlowDefinitionController(ctx context.Context, req *UpdateFlowDefinitionReq) (err error) {
	b := DB().FlowDefinition.UpdateOneID(req.ID)
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

func DeleteFlowDefinitionController(ctx context.Context, req *DeleteFlowDefinitionReq) (err error) {
	err = DB().FlowDefinition.DeleteOneID(req.ID).Exec(ctx)
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

func DeleteManyFlowDefinitionController(ctx context.Context, req *DeleteManyFlowDefinitionReq) (err error) {
	_, err = DB().FlowDefinition.Delete().Where(flowdefinition.IDIn(req.IDs...)).Exec(ctx)
	if err != nil {
		return gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	return
}

func CreateFlowDefinitionFormView() view.Node {
	return node.FormGrid("grid").MaxColumns(2).Children()
}

func ReadFlowDefinitionFormView() view.Node {
	return node.FormGrid("grid").MaxColumns(2).Children()
}

func UpdateFlowDefinitionFormView() view.Node {
	return node.FormGrid("grid").MaxColumns(2).Children()
}

func ListFlowDefinitionTableColumns() []view.Node {
	return []view.Node{}
}

func ListFlowDefinitionView() view.Node {
	return node.List("flowdefinitionList").
		AC(flowdefinitionViewACL["list:FlowDefinition"]).
		ForInit("@listFlowDefinition", ListFlowDefinitionController).
		DecoratorCard().
		EnableFilter().
		SelectionMultiple(true).
		Children(
			node.ListTable("flowdefinitionListTable").
				ActionBar(
					node.ListAction("flowdefinitionListActions").Children(
						node.ListActionRecordFormDrawer("addFlowDefinition").
							Title("addNew").
							AC(flowdefinitionViewACL["create:FlowDefinition"]).
							ButtonType("primary").
							ButtonIcon("IconPlus").
							ButtonPosition("left").
							Body(CreateFlowDefinitionFormView()).
							Footer(
								node.ActionFormDrawerCancel("cancel"),
								node.ActionFormDrawerSubmit("submit").ForSubmit("@addFlowDefinition", CreateFlowDefinitionController),
							),
						node.ListActionRowSelection("deleteManyFlowDefinition").
							Title("bulkDelete").
							AC(flowdefinitionViewACL["deleteMany:FlowDefinition"]).
							ForSubmit("@deleteManyFlowDefinition", DeleteManyFlowDefinitionController).
							AfterReload(true).
							ConfirmTitle("confirmDelete").
							ButtonStatus("danger").
							ButtonIcon("IconDelete").
							ButtonPosition("left"),
					),
				).
				Columns(
					append(
						ListFlowDefinitionTableColumns(),
						node.ListTableColumn("columnAction").
							Title("columnAction").
							DataIndex("__action").
							Width(150).
							Render(
								node.Space("actions").Size(0).SplitByDivider().Children(
									node.ListActionRecordFormDrawer("view").
										AC(flowdefinitionViewACL["get:FlowDefinition"]).
										ForInit("@getFlowDefinition", GetFlowDefinitionController).
										DrawerTitle("viewDrawerTitle").
										ButtonType("text").
										ButtonIcon("IconEye").
										Body(ReadFlowDefinitionFormView()).
										Footer(
											node.ActionFormDrawerCancel("cancel"),
										),
									node.ListActionRecordFormDrawer("edit").
										AC(flowdefinitionViewACL["update:FlowDefinition"]).
										ForInit("@getFlowDefinition", GetFlowDefinitionController).
										DrawerTitle("editDrawerTitle").
										ButtonType("text").
										ButtonIcon("IconEdit").
										Body(UpdateFlowDefinitionFormView()).
										Footer(
											node.ActionFormDrawerCancel("cancel"),
											node.ActionFormDrawerSubmit("submit").
												ForSubmit("@updateFlowDefinition", UpdateFlowDefinitionController),
										),
									node.ListActionRecordDelete("delete").
										AC(flowdefinitionViewACL["delete:FlowDefinition"]).
										ForSubmit("@deleteFlowDefinition", DeleteFlowDefinitionController).
										ButtonType("text").
										ButtonIcon("IconDelete"),
								),
							),
					)...,
				),
		)
}

var flowdefinitionViewACL = map[string]ac.AC{}

func init() {
	rf := reflect.ValueOf(schema.FlowDefinition{}).MethodByName("ACL")
	if rf.IsValid() {
		flowdefinitionViewACL, _ = rf.Call([]reflect.Value{})[0].Interface().(map[string]ac.AC)
	}
}

// FlowDefinitions is a parsable slice of FlowDefinition.
type FlowDefinitions []*FlowDefinition

func (fd FlowDefinitions) config(cfg config) {
	for _i := range fd {
		fd[_i].config = cfg
	}
}
