package ec2metadata

// GetInstanceID return ec2 instance id
func (c *Client) GetInstanceID() (string, error) {
	instanceID, err := c.svc.GetMetadata("instance-id")
	if err != nil {
		return "", err
	}
	return instanceID, nil
}

// GetRegion return region of ec2
func (c *Client) GetRegion() (string, error) {
	region, err := c.svc.GetMetadata("placement/region")
	if err != nil {
		return "", err
	}
	return region, nil
}
