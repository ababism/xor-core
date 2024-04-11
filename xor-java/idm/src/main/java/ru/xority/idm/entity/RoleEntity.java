package ru.xority.idm.entity;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.time.LocalDateTime;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class RoleEntity {
    public static final String UUID_FIELD = "uuid";
    public static final String NAME_FIELD = "name";
    public static final String CREATED_BY_UUID_FIELD = "created_by_uuid";
    public static final String CREATED_AT_FIELD = "created_at";
    public static final String ACTIVE_FIELD = "active";

    private UUID uuid;
    private String name;
    private UUID createdByUuid;
    private LocalDateTime createdAt;
    private boolean active;

    public static RoleEntity fromResultSet(ResultSet rs) throws SQLException {
        return new RoleEntity(
                UUID.fromString(rs.getString(UUID_FIELD)),
                rs.getString(NAME_FIELD),
                UUID.fromString(rs.getString(CREATED_BY_UUID_FIELD)),
                rs.getTimestamp(CREATED_AT_FIELD).toLocalDateTime(),
                rs.getBoolean(ACTIVE_FIELD)
        );
    }

    public static RoleEntity fromCreateRequest(String name, UUID createdByUuid) {
        return new RoleEntity(
                UUID.randomUUID(),
                name,
                createdByUuid,
                LocalDateTime.now(),
                true
        );
    }
}
