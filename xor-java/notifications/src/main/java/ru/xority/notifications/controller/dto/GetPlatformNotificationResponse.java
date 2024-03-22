package ru.xority.notifications.controller.dto;

import java.time.LocalDateTime;
import java.util.Optional;
import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

import ru.xority.notifications.entity.PlatformNotificationEntity;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
@JsonInclude(JsonInclude.Include.NON_EMPTY)
public class GetPlatformNotificationResponse {
    @JsonProperty("uuid")
    private UUID uuid;
    @JsonProperty("senderUuid")
    private UUID senderUuid;
    @JsonProperty("receiverUuid")
    private UUID receiverUuid;
    @JsonProperty("message")
    private String message;
    @JsonProperty("checked")
    private boolean checked;
    @JsonProperty("createdAt")
    private LocalDateTime createdAt;
    @JsonProperty("checkedAt")
    private Optional<LocalDateTime> checkedAt;

    public static GetPlatformNotificationResponse fromPlatformNotificationEntity(PlatformNotificationEntity platformNotification) {
        return new GetPlatformNotificationResponse(
                platformNotification.getUuid(),
                platformNotification.getSenderUuid(),
                platformNotification.getReceiverUuid(),
                platformNotification.getMessage(),
                platformNotification.isChecked(),
                platformNotification.getCreatedAt(),
                platformNotification.getCheckedAt()
        );
    }
}
