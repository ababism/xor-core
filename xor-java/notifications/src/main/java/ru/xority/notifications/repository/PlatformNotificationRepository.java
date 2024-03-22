package ru.xority.notifications.repository;

import java.util.List;
import java.util.Optional;

import ru.xority.notifications.entity.PlatformNotificationEntity;
import ru.xority.notifications.entity.PlatformNotificationFilter;

/**
 * @author foxleren
 */
public interface PlatformNotificationRepository {
    List<PlatformNotificationEntity> list(PlatformNotificationFilter filter);

    Optional<PlatformNotificationEntity> get(PlatformNotificationFilter filter);

    void create(PlatformNotificationEntity platformNotification);

    void update(PlatformNotificationEntity platformNotification);
}
