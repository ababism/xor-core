package ru.xority.idm.api.http.dto;

import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.NotNull;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
public class AssignRoleRequest {
    @JsonProperty("account_uuid")
    @NotNull
    private UUID accountUuid;
    @JsonProperty("role_uuid")
    @NotNull
    private UUID roleUuid;
}
