/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AccessAndMobilitySubscriptionData struct {
	SupportedFeatures           string                  `json:"supportedFeatures,omitempty" bson:"supportedFeatures" yaml:"supportedFeatures,omitempty"`
	Gpsis                       []string                `json:"gpsis,omitempty" bson:"gpsis" yaml:"gpsis,omitempty"`
	InternalGroupIds            []string                `json:"internalGroupIds,omitempty" bson:"internalGroupIds" yaml:"internalGroupIds,omitempty"`
	SubscribedUeAmbr            *AmbrRm                 `json:"subscribedUeAmbr,omitempty" bson:"subscribedUeAmbr" yaml:"subscribedUeAmbr,omitempty"`
	Nssai                       *Nssai                  `json:"nssai,omitempty" bson:"nssai" yaml:"nssai,omitempty"`
	RatRestrictions             []RatType               `json:"ratRestrictions,omitempty" bson:"ratRestrictions" yaml:"ratRestrictions,omitempty"`
	ForbiddenAreas              []Area                  `json:"forbiddenAreas,omitempty" bson:"forbiddenAreas" yaml:"forbiddenAreas,omitempty"`
	ServiceAreaRestriction      *ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty" bson:"serviceAreaRestriction" yaml:"serviceAreaRestriction,omitempty"`
	CoreNetworkTypeRestrictions []CoreNetworkType       `json:"coreNetworkTypeRestrictions,omitempty" bson:"coreNetworkTypeRestrictions" yaml:"coreNetworkTypeRestrictions,omitempty"`
	RfspIndex                   int32                   `json:"rfspIndex,omitempty" bson:"rfspIndex" yaml:"rfspIndex,omitempty"`
	SubsRegTimer                int32                   `json:"subsRegTimer,omitempty" bson:"subsRegTimer" yaml:"subsRegTimer,omitempty"`
	UeUsageType                 int32                   `json:"ueUsageType,omitempty" bson:"ueUsageType" yaml:"ueUsageType,omitempty"`
	MpsPriority                 bool                    `json:"mpsPriority,omitempty" bson:"mpsPriority" yaml:"mpsPriority,omitempty"`
	McsPriority                 bool                    `json:"mcsPriority,omitempty" bson:"mcsPriority" yaml:"mcsPriority,omitempty"`
	ActiveTime                  int32                   `json:"activeTime,omitempty" bson:"activeTime" yaml:"activeTime,omitempty"`
	DlPacketCount               int32                   `json:"dlPacketCount,omitempty" bson:"dlPacketCount" yaml:"dlPacketCount,omitempty"`
	SorInfo                     *SorInfo                `json:"sorInfo,omitempty" bson:"sorInfo" yaml:"sorInfo,omitempty"`
	MicoAllowed                 bool                    `json:"micoAllowed,omitempty" bson:"micoAllowed" yaml:"micoAllowed,omitempty"`
	SharedAmDataIds             []string                `json:"sharedAmDataIds,omitempty" bson:"sharedAmDataIds" yaml:"sharedAmDataIds,omitempty"`
	OdbPacketServices           OdbPacketServices       `json:"odbPacketServices,omitempty" bson:"odbPacketServices" yaml:"odbPacketServices,omitempty"`
}