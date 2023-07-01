package net

import "splendor/consts"

func (c *Client) GetRoomList() {
	c.SendMsg(consts.MessageType_Req_GetRoomList, []byte(""))
}
