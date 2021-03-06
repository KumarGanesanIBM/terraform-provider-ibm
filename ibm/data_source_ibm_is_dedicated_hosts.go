// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func dataSourceIbmIsDedicatedHosts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmIsDedicatedHostsRead,

		Schema: map[string]*schema.Schema{
			"host_group": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique identifier of the dedicated host group this dedicated host belongs to",
			},
			"dedicated_hosts": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of dedicated hosts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"available_memory": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The amount of memory in gibibytes that is currently available for instances.",
						},
						"available_vcpu": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The available VCPU for the dedicated host.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"architecture": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The VCPU architecture.",
									},
									"count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of VCPUs assigned.",
									},
								},
							},
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that the dedicated host was created.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this dedicated host.",
						},
						"host_group": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name of the dedicated host group this dedicated host is in.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this dedicated host.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this dedicated host.",
						},
						"instance_placement_enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, instances can be placed on this dedicated host.",
						},
						"instances": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Array of instances that are allocated to this dedicated host.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this virtual server instance.",
									},
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"more_info": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Link to documentation about deleted resources.",
												},
											},
										},
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this virtual server instance.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this virtual server instance.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this virtual server instance (and default system hostname).",
									},
								},
							},
						},
						"lifecycle_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The lifecycle state of the dedicated host resource.",
						},
						"memory": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total amount of memory in gibibytes for this host.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.",
						},
						"profile": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The profile this dedicated host uses.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this dedicated host.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this dedicated host profile.",
									},
								},
							},
						},
						"provisionable": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether this dedicated host is available for instance creation.",
						},
						"resource_group": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier of the resource group for this dedicated host.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of resource referenced.",
						},
						"socket_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of sockets for this host.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The administrative state of the dedicated host.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the dedicated host on which the unexpected property value was encountered.",
						},
						"supported_instance_profiles": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Array of instance profiles that can be used by instances placed on this dedicated host.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this virtual server instance profile.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this virtual server instance profile.",
									},
								},
							},
						},
						"vcpu": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The total VCPU of the dedicated host.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"architecture": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The VCPU architecture.",
									},
									"count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of VCPUs assigned.",
									},
								},
							},
						},
						"zone": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name of the zone this dedicated host resides in.",
						},
					},
				},
			},
			"first": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A link to the first page of resources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for a page of resources.",
						},
					},
				},
			},
			"limit": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum number of resources that can be returned by the request.",
			},
			"next": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A link to the next page of resources. This property is present for all pagesexcept the last page.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for a page of resources.",
						},
					},
				},
			},
			"total_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of resources across all pages.",
			},
		},
	}
}

func dataSourceIbmIsDedicatedHostsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	listDedicatedHostsOptions := &vpcv1.ListDedicatedHostsOptions{}
	if hostgroupintf, ok := d.GetOk("host_group"); ok {
		hostgroupid := hostgroupintf.(string)
		listDedicatedHostsOptions.DedicatedHostGroupID = &hostgroupid
	}

	dedicatedHostCollection, response, err := vpcClient.ListDedicatedHostsWithContext(context, listDedicatedHostsOptions)
	if err != nil {
		log.Printf("[DEBUG] ListDedicatedHostsWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	if len(dedicatedHostCollection.DedicatedHosts) != 0 {

		d.SetId(dataSourceIbmIsDedicatedHostsID(d))

		if dedicatedHostCollection.DedicatedHosts != nil {
			err = d.Set("dedicated_hosts", dataSourceDedicatedHostCollectionFlattenDedicatedHosts(dedicatedHostCollection.DedicatedHosts))
			if err != nil {
				return diag.FromErr(fmt.Errorf("Error setting dedicated_hosts %s", err))
			}
		}

		if dedicatedHostCollection.First != nil {
			err = d.Set("first", dataSourceDedicatedHostCollectionFlattenFirst(*dedicatedHostCollection.First))
			if err != nil {
				return diag.FromErr(fmt.Errorf("Error setting first %s", err))
			}
		}
		if err = d.Set("limit", dedicatedHostCollection.Limit); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting limit: %s", err))
		}

		if dedicatedHostCollection.Next != nil {
			err = d.Set("next", dataSourceDedicatedHostCollectionFlattenNext(*dedicatedHostCollection.Next))
			if err != nil {
				return diag.FromErr(fmt.Errorf("Error setting next %s", err))
			}
		}

		if err = d.Set("total_count", dedicatedHostCollection.TotalCount); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting total_count: %s", err))
		}
	}
	return nil
}

// dataSourceIbmIsDedicatedHostsID returns a reasonable ID for the list.
func dataSourceIbmIsDedicatedHostsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceDedicatedHostCollectionFlattenDedicatedHosts(result []vpcv1.DedicatedHost) (dedicatedHosts []map[string]interface{}) {
	for _, dedicatedHostsItem := range result {
		dedicatedHosts = append(dedicatedHosts, dataSourceDedicatedHostCollectionDedicatedHostsToMap(dedicatedHostsItem))
	}

	return dedicatedHosts
}

func dataSourceDedicatedHostCollectionDedicatedHostsToMap(dedicatedHostsItem vpcv1.DedicatedHost) (dedicatedHostsMap map[string]interface{}) {
	dedicatedHostsMap = map[string]interface{}{}

	if dedicatedHostsItem.AvailableMemory != nil {
		dedicatedHostsMap["available_memory"] = dedicatedHostsItem.AvailableMemory
	}
	if dedicatedHostsItem.AvailableVcpu != nil {
		availableVcpuList := []map[string]interface{}{}
		availableVcpuMap := dataSourceDedicatedHostCollectionDedicatedHostsAvailableVcpuToMap(*dedicatedHostsItem.AvailableVcpu)
		availableVcpuList = append(availableVcpuList, availableVcpuMap)
		dedicatedHostsMap["available_vcpu"] = availableVcpuList
	}
	if dedicatedHostsItem.CreatedAt != nil {
		dedicatedHostsMap["created_at"] = dedicatedHostsItem.CreatedAt.String()
	}
	if dedicatedHostsItem.CRN != nil {
		dedicatedHostsMap["crn"] = dedicatedHostsItem.CRN
	}
	if dedicatedHostsItem.Group != nil {
		dedicatedHostsMap["host_group"] = *dedicatedHostsItem.Group.ID
	}
	if dedicatedHostsItem.Href != nil {
		dedicatedHostsMap["href"] = dedicatedHostsItem.Href
	}
	if dedicatedHostsItem.ID != nil {
		dedicatedHostsMap["id"] = dedicatedHostsItem.ID
	}
	if dedicatedHostsItem.InstancePlacementEnabled != nil {
		dedicatedHostsMap["instance_placement_enabled"] = dedicatedHostsItem.InstancePlacementEnabled
	}
	if dedicatedHostsItem.Instances != nil {
		instancesList := []map[string]interface{}{}
		for _, instancesItem := range dedicatedHostsItem.Instances {
			instancesList = append(instancesList, dataSourceDedicatedHostCollectionDedicatedHostsInstancesToMap(instancesItem))
		}
		dedicatedHostsMap["instances"] = instancesList
	}
	if dedicatedHostsItem.LifecycleState != nil {
		dedicatedHostsMap["lifecycle_state"] = dedicatedHostsItem.LifecycleState
	}
	if dedicatedHostsItem.Memory != nil {
		dedicatedHostsMap["memory"] = dedicatedHostsItem.Memory
	}
	if dedicatedHostsItem.Name != nil {
		dedicatedHostsMap["name"] = dedicatedHostsItem.Name
	}
	if dedicatedHostsItem.Profile != nil {
		profileList := []map[string]interface{}{}
		profileMap := dataSourceDedicatedHostCollectionDedicatedHostsProfileToMap(*dedicatedHostsItem.Profile)
		profileList = append(profileList, profileMap)
		dedicatedHostsMap["profile"] = profileList
	}
	if dedicatedHostsItem.Provisionable != nil {
		dedicatedHostsMap["provisionable"] = dedicatedHostsItem.Provisionable
	}
	if dedicatedHostsItem.ResourceGroup != nil {
		dedicatedHostsMap["resource_group"] = *dedicatedHostsItem.ResourceGroup.ID
	}
	if dedicatedHostsItem.ResourceType != nil {
		dedicatedHostsMap["resource_type"] = dedicatedHostsItem.ResourceType
	}
	if dedicatedHostsItem.SocketCount != nil {
		dedicatedHostsMap["socket_count"] = dedicatedHostsItem.SocketCount
	}
	if dedicatedHostsItem.State != nil {
		dedicatedHostsMap["state"] = dedicatedHostsItem.State
	}
	if dedicatedHostsItem.SupportedInstanceProfiles != nil {
		supportedInstanceProfilesList := []map[string]interface{}{}
		for _, supportedInstanceProfilesItem := range dedicatedHostsItem.SupportedInstanceProfiles {
			supportedInstanceProfilesList = append(supportedInstanceProfilesList, dataSourceDedicatedHostCollectionDedicatedHostsSupportedInstanceProfilesToMap(supportedInstanceProfilesItem))
		}
		dedicatedHostsMap["supported_instance_profiles"] = supportedInstanceProfilesList
	}
	if dedicatedHostsItem.Vcpu != nil {
		vcpuList := []map[string]interface{}{}
		vcpuMap := dataSourceDedicatedHostCollectionDedicatedHostsVcpuToMap(*dedicatedHostsItem.Vcpu)
		vcpuList = append(vcpuList, vcpuMap)
		dedicatedHostsMap["vcpu"] = vcpuList
	}
	if dedicatedHostsItem.Zone != nil {
		dedicatedHostsMap["zone"] = *dedicatedHostsItem.Zone.Name
	}

	return dedicatedHostsMap
}

func dataSourceDedicatedHostCollectionDedicatedHostsAvailableVcpuToMap(availableVcpuItem vpcv1.Vcpu) (availableVcpuMap map[string]interface{}) {
	availableVcpuMap = map[string]interface{}{}

	if availableVcpuItem.Architecture != nil {
		availableVcpuMap["architecture"] = availableVcpuItem.Architecture
	}
	if availableVcpuItem.Count != nil {
		availableVcpuMap["count"] = availableVcpuItem.Count
	}

	return availableVcpuMap
}

func dataSourceDedicatedHostCollectionDedicatedHostsGroupToMap(groupItem vpcv1.DedicatedHostGroupReference) (groupMap map[string]interface{}) {
	groupMap = map[string]interface{}{}

	if groupItem.CRN != nil {
		groupMap["crn"] = groupItem.CRN
	}
	if groupItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceDedicatedHostCollectionGroupDeletedToMap(*groupItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		groupMap["deleted"] = deletedList
	}
	if groupItem.Href != nil {
		groupMap["href"] = groupItem.Href
	}
	if groupItem.ID != nil {
		groupMap["id"] = groupItem.ID
	}
	if groupItem.Name != nil {
		groupMap["name"] = groupItem.Name
	}
	if groupItem.ResourceType != nil {
		groupMap["resource_type"] = groupItem.ResourceType
	}

	return groupMap
}

func dataSourceDedicatedHostCollectionGroupDeletedToMap(deletedItem vpcv1.DedicatedHostGroupReferenceDeleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceDedicatedHostCollectionInstancesDeletedToMap(deletedItem vpcv1.InstanceReferenceDeleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceDedicatedHostCollectionDedicatedHostsInstancesToMap(instancesItem vpcv1.InstanceReference) (instancesMap map[string]interface{}) {
	instancesMap = map[string]interface{}{}

	if instancesItem.CRN != nil {
		instancesMap["crn"] = instancesItem.CRN
	}
	if instancesItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceDedicatedHostCollectionInstancesDeletedToMap(*instancesItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		instancesMap["deleted"] = deletedList
	}
	if instancesItem.Href != nil {
		instancesMap["href"] = instancesItem.Href
	}
	if instancesItem.ID != nil {
		instancesMap["id"] = instancesItem.ID
	}
	if instancesItem.Name != nil {
		instancesMap["name"] = instancesItem.Name
	}

	return instancesMap
}

func dataSourceDedicatedHostCollectionDedicatedHostsProfileToMap(profileItem vpcv1.DedicatedHostProfileReference) (profileMap map[string]interface{}) {
	profileMap = map[string]interface{}{}

	if profileItem.Href != nil {
		profileMap["href"] = profileItem.Href
	}
	if profileItem.Name != nil {
		profileMap["name"] = profileItem.Name
	}

	return profileMap
}

func dataSourceDedicatedHostCollectionDedicatedHostsSupportedInstanceProfilesToMap(supportedInstanceProfilesItem vpcv1.InstanceProfileReference) (supportedInstanceProfilesMap map[string]interface{}) {
	supportedInstanceProfilesMap = map[string]interface{}{}

	if supportedInstanceProfilesItem.Href != nil {
		supportedInstanceProfilesMap["href"] = supportedInstanceProfilesItem.Href
	}
	if supportedInstanceProfilesItem.Name != nil {
		supportedInstanceProfilesMap["name"] = supportedInstanceProfilesItem.Name
	}

	return supportedInstanceProfilesMap
}

func dataSourceDedicatedHostCollectionDedicatedHostsVcpuToMap(vcpuItem vpcv1.Vcpu) (vcpuMap map[string]interface{}) {
	vcpuMap = map[string]interface{}{}

	if vcpuItem.Architecture != nil {
		vcpuMap["architecture"] = vcpuItem.Architecture
	}
	if vcpuItem.Count != nil {
		vcpuMap["count"] = vcpuItem.Count
	}

	return vcpuMap
}

func dataSourceDedicatedHostCollectionFlattenFirst(result vpcv1.DedicatedHostCollectionFirst) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceDedicatedHostCollectionFirstToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceDedicatedHostCollectionFirstToMap(firstItem vpcv1.DedicatedHostCollectionFirst) (firstMap map[string]interface{}) {
	firstMap = map[string]interface{}{}

	if firstItem.Href != nil {
		firstMap["href"] = firstItem.Href
	}

	return firstMap
}

func dataSourceDedicatedHostCollectionFlattenNext(result vpcv1.DedicatedHostCollectionNext) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceDedicatedHostCollectionNextToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceDedicatedHostCollectionNextToMap(nextItem vpcv1.DedicatedHostCollectionNext) (nextMap map[string]interface{}) {
	nextMap = map[string]interface{}{}

	if nextItem.Href != nil {
		nextMap["href"] = nextItem.Href
	}

	return nextMap
}
