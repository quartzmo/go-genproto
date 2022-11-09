// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package essentialcontacts aliases all exported identifiers in package
// "cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb".
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package essentialcontacts

import (
	src "cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
const (
	NotificationCategory_ALL                               = src.NotificationCategory_ALL
	NotificationCategory_BILLING                           = src.NotificationCategory_BILLING
	NotificationCategory_LEGAL                             = src.NotificationCategory_LEGAL
	NotificationCategory_NOTIFICATION_CATEGORY_UNSPECIFIED = src.NotificationCategory_NOTIFICATION_CATEGORY_UNSPECIFIED
	NotificationCategory_PRODUCT_UPDATES                   = src.NotificationCategory_PRODUCT_UPDATES
	NotificationCategory_SECURITY                          = src.NotificationCategory_SECURITY
	NotificationCategory_SUSPENSION                        = src.NotificationCategory_SUSPENSION
	NotificationCategory_TECHNICAL                         = src.NotificationCategory_TECHNICAL
	NotificationCategory_TECHNICAL_INCIDENTS               = src.NotificationCategory_TECHNICAL_INCIDENTS
	ValidationState_INVALID                                = src.ValidationState_INVALID
	ValidationState_VALID                                  = src.ValidationState_VALID
	ValidationState_VALIDATION_STATE_UNSPECIFIED           = src.ValidationState_VALIDATION_STATE_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
var (
	File_google_cloud_essentialcontacts_v1_enums_proto   = src.File_google_cloud_essentialcontacts_v1_enums_proto
	File_google_cloud_essentialcontacts_v1_service_proto = src.File_google_cloud_essentialcontacts_v1_service_proto
	NotificationCategory_name                            = src.NotificationCategory_name
	NotificationCategory_value                           = src.NotificationCategory_value
	ValidationState_name                                 = src.ValidationState_name
	ValidationState_value                                = src.ValidationState_value
)

// Request message for the ComputeContacts method.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type ComputeContactsRequest = src.ComputeContactsRequest

// Response message for the ComputeContacts method.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type ComputeContactsResponse = src.ComputeContactsResponse

// A contact that will receive notifications from Google Cloud.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type Contact = src.Contact

// Request message for the CreateContact method.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type CreateContactRequest = src.CreateContactRequest

// Request message for the DeleteContact method.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type DeleteContactRequest = src.DeleteContactRequest

// EssentialContactsServiceClient is the client API for
// EssentialContactsService service. For semantics around ctx use and
// closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type EssentialContactsServiceClient = src.EssentialContactsServiceClient

// EssentialContactsServiceServer is the server API for
// EssentialContactsService service.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type EssentialContactsServiceServer = src.EssentialContactsServiceServer

// Request message for the GetContact method.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type GetContactRequest = src.GetContactRequest

// Request message for the ListContacts method.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type ListContactsRequest = src.ListContactsRequest

// Response message for the ListContacts method.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type ListContactsResponse = src.ListContactsResponse

// The notification categories that an essential contact can be subscribed to.
// Each notification will be categorized by the sender into one of the
// following categories. All contacts that are subscribed to that category will
// receive the notification.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type NotificationCategory = src.NotificationCategory

// Request message for the SendTestMessage method.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type SendTestMessageRequest = src.SendTestMessageRequest

// UnimplementedEssentialContactsServiceServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type UnimplementedEssentialContactsServiceServer = src.UnimplementedEssentialContactsServiceServer

// Request message for the UpdateContact method.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type UpdateContactRequest = src.UpdateContactRequest

// A contact's validation state indicates whether or not it is the correct
// contact to be receiving notifications for a particular resource.
//
// Deprecated: Please use types in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
type ValidationState = src.ValidationState

// Deprecated: Please use funcs in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
func NewEssentialContactsServiceClient(cc grpc.ClientConnInterface) EssentialContactsServiceClient {
	return src.NewEssentialContactsServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb
func RegisterEssentialContactsServiceServer(s *grpc.Server, srv EssentialContactsServiceServer) {
	src.RegisterEssentialContactsServiceServer(s, srv)
}