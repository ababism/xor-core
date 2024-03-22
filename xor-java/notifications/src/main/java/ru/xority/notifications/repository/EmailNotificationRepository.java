package ru.xority.notifications.repository;

import java.util.List;
import java.util.Optional;

import ru.xority.notifications.entity.EmailNotificationEntity;
import ru.xority.notifications.entity.EmailNotificationFilter;

/**
 * @author foxleren
 */

public interface EmailNotificationRepository {
    List<EmailNotificationEntity> list(EmailNotificationFilter filter);

    Optional<EmailNotificationEntity> get(EmailNotificationFilter filter);

    void create(EmailNotificationEntity notification);
}
