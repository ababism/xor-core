package ru.xority.notifications.service;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import ru.xority.exception.BadRequestException;
import ru.xority.notifications.entity.PlatformNotificationEntity;
import ru.xority.notifications.entity.PlatformNotificationFilter;
import ru.xority.notifications.repository.PlatformNotificationRepository;

/**
 * @author foxleren
 */
@Service
@RequiredArgsConstructor
public class PlatformNotificationService {
    private final PlatformNotificationRepository platformNotificationRepository;

    public List<PlatformNotificationEntity> list(PlatformNotificationFilter filter) {
        return platformNotificationRepository.list(filter);
    }

    public String create(PlatformNotificationEntity platformNotification) {
        return platformNotificationRepository.create(platformNotification);
    }

    public void check(String notificationId) {
        PlatformNotificationFilter filter = PlatformNotificationFilter.byId(notificationId);
        Optional<PlatformNotificationEntity> platformNotificationO = platformNotificationRepository.get(filter);

        platformNotificationO.orElseThrow(() -> new BadRequestException("Platform notification is not found"));
        PlatformNotificationEntity platformNotification = platformNotificationO.get();

        if (platformNotification.isChecked()) {
            throw new BadRequestException("Platform notification is already checked");
        }

        platformNotification.setChecked(true);
        platformNotification.setCheckedAt(Optional.of(LocalDateTime.now()));

        platformNotificationRepository.update(platformNotification);
    }
}
