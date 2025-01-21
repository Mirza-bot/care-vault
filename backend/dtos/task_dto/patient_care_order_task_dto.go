package taskdto

import userdtos "care-vault/dtos/user_dtos"

type CareOrderTaskDto struct {
    BaseTaskDto
    Patient userdtos.BaseUserDto 
}
