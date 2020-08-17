/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AssignedEbiData struct {
	PduSessionId    int32           `json:"pduSessionId"`
	AssignedEbiList []EbiArpMapping `json:"assignedEbiList,omitempty"`
	FailedArpList   []Arp           `json:"failedArpList,omitempty"`
	ReleasedEbiList []int32         `json:"releasedEbiList,omitempty"`
}