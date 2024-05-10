package service

import "seckill/model"

var Activity model.Activities

func CreateActivities(adminid int) {
	Activity = model.Activities{
		AdminId: adminid,
	}
}
