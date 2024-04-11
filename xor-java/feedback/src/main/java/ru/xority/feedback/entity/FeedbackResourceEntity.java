package ru.xority.feedback.entity;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.time.LocalDateTime;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

import ru.xority.feedback.controller.dto.CreateFeedbackResourceRequest;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class FeedbackResourceEntity {
    public static final String UUID_FIELD = "uuid";
    public static final String NAME_FIELD = "name";
    public static final String DESCRIPTION_FIELD = "description";
    public static final String CREATED_BY_UUID_FIELD = "created_by_uuid";
    public static final String CREATED_AT_FIELD = "created_at";
    public static final String ACTIVE_FIELD = "active";


    private UUID uuid;
    private String name;
    private String description;
    private UUID createdByUuid;
    private LocalDateTime createdAt;
    private boolean active;

    public static FeedbackResourceEntity fromResultSet(ResultSet rs) throws SQLException {
        return new FeedbackResourceEntity(
                UUID.fromString(rs.getString(UUID_FIELD)),
                rs.getString(NAME_FIELD),
                rs.getString(DESCRIPTION_FIELD),
                UUID.fromString(rs.getString(CREATED_BY_UUID_FIELD)),
                rs.getTimestamp(CREATED_AT_FIELD).toLocalDateTime(),
                rs.getBoolean(ACTIVE_FIELD)
        );
    }

    public static FeedbackResourceEntity fromCreateFeedbackResourceRequest(UUID createdByUuid, CreateFeedbackResourceRequest request) {
        return new FeedbackResourceEntity(
                UUID.randomUUID(),
                request.getName(),
                request.getDescription(),
                createdByUuid,
                LocalDateTime.now(),
                true
        );
    }
}
