package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct QosFlowAddOrModifyResponseList */
/* QosFlowAddOrModifyResponseItem */
type QosFlowAddOrModifyResponseList struct {
	List []QosFlowAddOrModifyResponseItem `aper:"valueExt,sizeLB:1,sizeUB:64"`
}