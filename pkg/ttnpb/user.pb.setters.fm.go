// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"
	time "time"

	types "github.com/gogo/protobuf/types"
)

func (dst *User) SetFields(src *User, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				newDst := &dst.UserIdentifiers
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIdentifiers = src.UserIdentifiers
				} else {
					var zero UserIdentifiers
					dst.UserIdentifiers = zero
				}
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "description":
			if len(subs) > 0 {
				return fmt.Errorf("'description' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Description = src.Description
			} else {
				var zero string
				dst.Description = zero
			}
		case "attributes":
			if len(subs) > 0 {
				return fmt.Errorf("'attributes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Attributes = src.Attributes
			} else {
				dst.Attributes = nil
			}
		case "contact_info":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_info' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactInfo = src.ContactInfo
			} else {
				dst.ContactInfo = nil
			}
		case "primary_email_address":
			if len(subs) > 0 {
				return fmt.Errorf("'primary_email_address' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PrimaryEmailAddress = src.PrimaryEmailAddress
			} else {
				var zero string
				dst.PrimaryEmailAddress = zero
			}
		case "primary_email_address_validated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'primary_email_address_validated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PrimaryEmailAddressValidatedAt = src.PrimaryEmailAddressValidatedAt
			} else {
				dst.PrimaryEmailAddressValidatedAt = nil
			}
		case "password":
			if len(subs) > 0 {
				return fmt.Errorf("'password' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Password = src.Password
			} else {
				var zero string
				dst.Password = zero
			}
		case "password_updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'password_updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PasswordUpdatedAt = src.PasswordUpdatedAt
			} else {
				var zero time.Time
				dst.PasswordUpdatedAt = zero
			}
		case "require_password_update":
			if len(subs) > 0 {
				return fmt.Errorf("'require_password_update' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.RequirePasswordUpdate = src.RequirePasswordUpdate
			} else {
				var zero bool
				dst.RequirePasswordUpdate = zero
			}
		case "state":
			if len(subs) > 0 {
				return fmt.Errorf("'state' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.State = src.State
			} else {
				var zero State
				dst.State = zero
			}
		case "admin":
			if len(subs) > 0 {
				return fmt.Errorf("'admin' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Admin = src.Admin
			} else {
				var zero bool
				dst.Admin = zero
			}
		case "temporary_password":
			if len(subs) > 0 {
				return fmt.Errorf("'temporary_password' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TemporaryPassword = src.TemporaryPassword
			} else {
				var zero string
				dst.TemporaryPassword = zero
			}
		case "temporary_password_created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'temporary_password_created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TemporaryPasswordCreatedAt = src.TemporaryPasswordCreatedAt
			} else {
				dst.TemporaryPasswordCreatedAt = nil
			}
		case "temporary_password_expires_at":
			if len(subs) > 0 {
				return fmt.Errorf("'temporary_password_expires_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TemporaryPasswordExpiresAt = src.TemporaryPasswordExpiresAt
			} else {
				dst.TemporaryPasswordExpiresAt = nil
			}
		case "profile_picture":
			if len(subs) > 0 {
				newDst := dst.ProfilePicture
				if newDst == nil {
					newDst = &Picture{}
					dst.ProfilePicture = newDst
				}
				var newSrc *Picture
				if src != nil {
					newSrc = src.ProfilePicture
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ProfilePicture = src.ProfilePicture
				} else {
					dst.ProfilePicture = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Picture) SetFields(src *Picture, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "embedded":
			if len(subs) > 0 {
				newDst := dst.Embedded
				if newDst == nil {
					newDst = &Picture_Embedded{}
					dst.Embedded = newDst
				}
				var newSrc *Picture_Embedded
				if src != nil {
					newSrc = src.Embedded
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Embedded = src.Embedded
				} else {
					dst.Embedded = nil
				}
			}
		case "sizes":
			if len(subs) > 0 {
				return fmt.Errorf("'sizes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Sizes = src.Sizes
			} else {
				dst.Sizes = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Users) SetFields(src *Users, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "users":
			if len(subs) > 0 {
				return fmt.Errorf("'users' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Users = src.Users
			} else {
				dst.Users = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetUserRequest) SetFields(src *GetUserRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIdentifiers
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIdentifiers = src.UserIdentifiers
				} else {
					var zero UserIdentifiers
					dst.UserIdentifiers = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateUserRequest) SetFields(src *CreateUserRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "user":
			if len(subs) > 0 {
				newDst := &dst.User
				var newSrc *User
				if src != nil {
					newSrc = &src.User
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.User = src.User
				} else {
					var zero User
					dst.User = zero
				}
			}
		case "invitation_token":
			if len(subs) > 0 {
				return fmt.Errorf("'invitation_token' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.InvitationToken = src.InvitationToken
			} else {
				var zero string
				dst.InvitationToken = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateUserRequest) SetFields(src *UpdateUserRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "user":
			if len(subs) > 0 {
				newDst := &dst.User
				var newSrc *User
				if src != nil {
					newSrc = &src.User
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.User = src.User
				} else {
					var zero User
					dst.User = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateTemporaryPasswordRequest) SetFields(src *CreateTemporaryPasswordRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIdentifiers
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIdentifiers = src.UserIdentifiers
				} else {
					var zero UserIdentifiers
					dst.UserIdentifiers = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateUserPasswordRequest) SetFields(src *UpdateUserPasswordRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIdentifiers
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIdentifiers = src.UserIdentifiers
				} else {
					var zero UserIdentifiers
					dst.UserIdentifiers = zero
				}
			}
		case "new":
			if len(subs) > 0 {
				return fmt.Errorf("'new' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.New = src.New
			} else {
				var zero string
				dst.New = zero
			}
		case "old":
			if len(subs) > 0 {
				return fmt.Errorf("'old' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Old = src.Old
			} else {
				var zero string
				dst.Old = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListUserAPIKeysRequest) SetFields(src *ListUserAPIKeysRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIdentifiers
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIdentifiers = src.UserIdentifiers
				} else {
					var zero UserIdentifiers
					dst.UserIdentifiers = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateUserAPIKeyRequest) SetFields(src *CreateUserAPIKeyRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIdentifiers
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIdentifiers = src.UserIdentifiers
				} else {
					var zero UserIdentifiers
					dst.UserIdentifiers = zero
				}
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "rights":
			if len(subs) > 0 {
				return fmt.Errorf("'rights' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Rights = src.Rights
			} else {
				dst.Rights = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateUserAPIKeyRequest) SetFields(src *UpdateUserAPIKeyRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIdentifiers
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIdentifiers = src.UserIdentifiers
				} else {
					var zero UserIdentifiers
					dst.UserIdentifiers = zero
				}
			}
		case "api_key":
			if len(subs) > 0 {
				newDst := &dst.APIKey
				var newSrc *APIKey
				if src != nil {
					newSrc = &src.APIKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.APIKey = src.APIKey
				} else {
					var zero APIKey
					dst.APIKey = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Invitation) SetFields(src *Invitation, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "email":
			if len(subs) > 0 {
				return fmt.Errorf("'email' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Email = src.Email
			} else {
				var zero string
				dst.Email = zero
			}
		case "token":
			if len(subs) > 0 {
				return fmt.Errorf("'token' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Token = src.Token
			} else {
				var zero string
				dst.Token = zero
			}
		case "expires_at":
			if len(subs) > 0 {
				return fmt.Errorf("'expires_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ExpiresAt = src.ExpiresAt
			} else {
				var zero time.Time
				dst.ExpiresAt = zero
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "accepted_at":
			if len(subs) > 0 {
				return fmt.Errorf("'accepted_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.AcceptedAt = src.AcceptedAt
			} else {
				dst.AcceptedAt = nil
			}
		case "accepted_by":
			if len(subs) > 0 {
				newDst := dst.AcceptedBy
				if newDst == nil {
					newDst = &UserIdentifiers{}
					dst.AcceptedBy = newDst
				}
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = src.AcceptedBy
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.AcceptedBy = src.AcceptedBy
				} else {
					dst.AcceptedBy = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListInvitationsRequest) SetFields(src *ListInvitationsRequest, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message ListInvitationsRequest has no fields, but paths %s were specified", paths)
	}
	if src != nil {
		*dst = *src
	}
	return nil
}

func (dst *Invitations) SetFields(src *Invitations, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "invitations":
			if len(subs) > 0 {
				return fmt.Errorf("'invitations' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Invitations = src.Invitations
			} else {
				dst.Invitations = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *SendInvitationRequest) SetFields(src *SendInvitationRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "email":
			if len(subs) > 0 {
				return fmt.Errorf("'email' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Email = src.Email
			} else {
				var zero string
				dst.Email = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeleteInvitationRequest) SetFields(src *DeleteInvitationRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "email":
			if len(subs) > 0 {
				return fmt.Errorf("'email' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Email = src.Email
			} else {
				var zero string
				dst.Email = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UserSessionIdentifiers) SetFields(src *UserSessionIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIdentifiers
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIdentifiers = src.UserIdentifiers
				} else {
					var zero UserIdentifiers
					dst.UserIdentifiers = zero
				}
			}
		case "session_id":
			if len(subs) > 0 {
				return fmt.Errorf("'session_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SessionID = src.SessionID
			} else {
				var zero string
				dst.SessionID = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UserSession) SetFields(src *UserSession, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIdentifiers
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIdentifiers = src.UserIdentifiers
				} else {
					var zero UserIdentifiers
					dst.UserIdentifiers = zero
				}
			}
		case "session_id":
			if len(subs) > 0 {
				return fmt.Errorf("'session_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SessionID = src.SessionID
			} else {
				var zero string
				dst.SessionID = zero
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "expires_at":
			if len(subs) > 0 {
				return fmt.Errorf("'expires_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ExpiresAt = src.ExpiresAt
			} else {
				dst.ExpiresAt = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UserSessions) SetFields(src *UserSessions, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "sessions":
			if len(subs) > 0 {
				return fmt.Errorf("'sessions' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Sessions = src.Sessions
			} else {
				dst.Sessions = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListUserSessionsRequest) SetFields(src *ListUserSessionsRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIdentifiers
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIdentifiers = src.UserIdentifiers
				} else {
					var zero UserIdentifiers
					dst.UserIdentifiers = zero
				}
			}
		case "order":
			if len(subs) > 0 {
				return fmt.Errorf("'order' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Order = src.Order
			} else {
				var zero string
				dst.Order = zero
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Picture_Embedded) SetFields(src *Picture_Embedded, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "mime_type":
			if len(subs) > 0 {
				return fmt.Errorf("'mime_type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MimeType = src.MimeType
			} else {
				var zero string
				dst.MimeType = zero
			}
		case "data":
			if len(subs) > 0 {
				return fmt.Errorf("'data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Data = src.Data
			} else {
				dst.Data = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
