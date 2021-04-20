/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type RemoveAccessControlGroupRuleParameter struct {

	// IP블록
IpBlock *string `json:"ipBlock,omitempty"`

	// 접근소스ACG
AccessControlGroupSequence *string `json:"accessControlGroupSequence,omitempty"`

	// 포트범위
PortRange *string `json:"portRange,omitempty"`

	// 프로토콜유형코드
ProtocolTypeCode *string `json:"protocolTypeCode"`
}