package ru.xority.feedback.entity;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.time.LocalDateTime;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

import ru.xority.feedback.controller.dto.CreateFeedbackItemRequest;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class FeedbackItemEntity {
    public static final String UUID_FIELD = "uuid";
    public static final String RESOURCE_UUID_FIELD = "resource_uuid";
    public static final String CREATED_BY_UUID_FIELD = "created_by_uuid";
    public static final String CREATED_AT_FIELD = "created_at";
    public static final String TEXT_FIELD = "text";
    public static final String RATING_FIELD = "rating";
    public static final String ACTIVE_FIELD = "active";

    private UUID uuid;
    private UUID resourceUuid;
    private UUID createdByUuid;
    private LocalDateTime createdAt;
    private String text;
    private int rating;
    private boolean active;

    public static FeedbackItemEntity fromResultSet(ResultSet rs) throws SQLException {
        return new FeedbackItemEntity(
                UUID.fromString(rs.getString(UUID_FIELD)),
                UUID.fromString(rs.getString(RESOURCE_UUID_FIELD)),
                UUID.fromString(rs.getString(CREATED_BY_UUID_FIELD)),
                rs.getTimestamp(CREATED_AT_FIELD).toLocalDateTime(),
                rs.getString(TEXT_FIELD),
                rs.getInt(RATING_FIELD),
                rs.getBoolean(ACTIVE_FIELD)
        );
    }

    public static FeedbackItemEntity fromCreateFeedbackItemRequest(UUID createdByUuid, CreateFeedbackItemRequest request) {
        return new FeedbackItemEntity(
                UUID.randomUUID(),
                request.getResourceUuid(),
                createdByUuid,
                LocalDateTime.now(),
                request.getText(),
                request.getRating(),
                true
        );
    }
}
