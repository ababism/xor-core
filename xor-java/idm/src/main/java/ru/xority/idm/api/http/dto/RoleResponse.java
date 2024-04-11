package ru.xority.idm.api.http.dto;

import java.time.LocalDateTime;
import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

import ru.xority.idm.entity.RoleEntity;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class RoleResponse {
    @JsonProperty("uuid")
    private UUID uuid;
    @JsonProperty("name")
    private String name;
    @JsonProperty("created_by_uuid")
    private UUID createdByUuid;
    @JsonProperty("created_at")
    private LocalDateTime createdAt;
    @JsonProperty("active")
    private boolean active;

    public static RoleResponse fromRoleEntity(RoleEntity role) {
        return new RoleResponse(
                role.getUuid(),
                role.getName(),
                role.getCreatedByUuid(),
                role.getCreatedAt(),
                role.isActive()
        );
    }
}
