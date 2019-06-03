# ConfigurationsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | [***StringConfigItem**](StringConfigItem.md) | The auth mode of current system, such as \&quot;db_auth\&quot;, \&quot;ldap_auth\&quot; | [optional] [default to null]
**EmailFrom** | [***StringConfigItem**](StringConfigItem.md) | The sender name for Email notification. | [optional] [default to null]
**EmailHost** | [***StringConfigItem**](StringConfigItem.md) | The hostname of SMTP server that sends Email notification. | [optional] [default to null]
**EmailPort** | [***IntegerConfigItem**](IntegerConfigItem.md) | The port of SMTP server. | [optional] [default to null]
**EmailIdentity** | [***StringConfigItem**](StringConfigItem.md) | By default it&#39;s empty so the email_username is picked. | [optional] [default to null]
**EmailUsername** | [***StringConfigItem**](StringConfigItem.md) | The username for authenticate against SMTP server. | [optional] [default to null]
**EmailSsl** | [***BoolConfigItem**](BoolConfigItem.md) | When it&#39;s set to true the system will access Email server via TLS by default.  If it&#39;s set to false, it still will handle \&quot;STARTTLS\&quot; from server side. | [optional] [default to null]
**EmailInsecure** | [***BoolConfigItem**](BoolConfigItem.md) | Whether or not the certificate will be verified when Harbor tries to access the email server. | [optional] [default to null]
**LdapUrl** | [***StringConfigItem**](StringConfigItem.md) | The URL of LDAP server. | [optional] [default to null]
**LdapBaseDn** | [***StringConfigItem**](StringConfigItem.md) | The Base DN for LDAP binding. | [optional] [default to null]
**LdapFilter** | [***StringConfigItem**](StringConfigItem.md) | The filter for LDAP binding. | [optional] [default to null]
**LdapScope** | **int32** | 0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE | [optional] [default to null]
**LdapUid** | [***StringConfigItem**](StringConfigItem.md) | The attribute which is used as identity for the LDAP binding, such as \&quot;CN\&quot; or \&quot;SAMAccountname\&quot; | [optional] [default to null]
**LdapSearchDn** | **string** | The DN of the user to do the search. | [optional] [default to null]
**LdapTimeout** | [***IntegerConfigItem**](IntegerConfigItem.md) | timeout in seconds for connection to LDAP server. | [optional] [default to null]
**LdapGroupAttributeName** | [***StringConfigItem**](StringConfigItem.md) | The attribute which is used as identity of the LDAP group, default is cn. | [optional] [default to null]
**LdapGroupBaseDn** | [***StringConfigItem**](StringConfigItem.md) | The base DN to search LDAP group. | [optional] [default to null]
**LdapGroupSearchFilter** | [***StringConfigItem**](StringConfigItem.md) | The filter to search the ldap group. | [optional] [default to null]
**LdapGroupSearchScope** | [***IntegerConfigItem**](IntegerConfigItem.md) | The scope to search ldap. &#39;0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE&#39; | [optional] [default to null]
**LdapGroupAdminDn** | [***StringConfigItem**](StringConfigItem.md) | Specify the ldap group which have the same privilege with Harbor admin. | [optional] [default to null]
**ProjectCreationRestriction** | [***StringConfigItem**](StringConfigItem.md) | This attribute restricts what users have the permission to create project.  It can be \&quot;everyone\&quot; or \&quot;adminonly\&quot;. | [optional] [default to null]
**ReadOnly** | [***BoolConfigItem**](BoolConfigItem.md) | &#39;docker push&#39; is prohibited by Harbor if you set it to true.    | [optional] [default to null]
**SelfRegistration** | [***BoolConfigItem**](BoolConfigItem.md) | Whether the Harbor instance supports self-registration.  If it&#39;s set to false, admin need to add user to the instance. | [optional] [default to null]
**TokenExpiration** | [***IntegerConfigItem**](IntegerConfigItem.md) | The expiration time of the token for internal Registry, in minutes. | [optional] [default to null]
**VerifyRemoteCert** | [***BoolConfigItem**](BoolConfigItem.md) | Whether or not the certificate will be verified when Harbor tries to access a remote Harbor instance for replication. | [optional] [default to null]
**ScanAllPolicy** | [***ConfigurationsScanAllPolicy**](Configurations_scan_all_policy.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


