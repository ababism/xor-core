package ru.xority.idm.mapper;

import org.mapstruct.Mapper;

import ru.xority.idm.controller.dto.AccountResponse;
import ru.xority.idm.entity.AccountEntity;

/**
 * @author foxleren
 */
@Mapper(componentModel = "spring")
public interface AccountMapper {
    AccountResponse accountEntityToAccountResponse(AccountEntity entity);
}