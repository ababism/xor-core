package ru.xority.notifications.controller;

import java.util.List;
import java.util.UUID;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import ru.xority.notifications.controller.dto.CreatePlatformNotificationRequest;
import ru.xority.notifications.controller.dto.CreatePlatformNotificationResponse;
import ru.xority.notifications.controller.dto.GetPlatformNotificationResponse;
import ru.xority.notifications.entity.PlatformNotificationEntity;
import ru.xority.notifications.entity.PlatformNotificationFilter;
import ru.xority.notifications.service.PlatformNotificationService;
import ru.xority.response.SuccessResponse;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/platform-notification")
@RequiredArgsConstructor
public class PlatformNotificationController {
    private final PlatformNotificationService platformNotificationService;

    @GetMapping("/list")
    public ResponseEntity<?> list(@RequestParam UUID recieverUuid) {
        PlatformNotificationFilter filter = PlatformNotificationFilter.byReceiverUuid(recieverUuid);
        List<GetPlatformNotificationResponse> platformNotifications = platformNotificationService.list(filter)
                .stream().map(GetPlatformNotificationResponse::fromPlatformNotificationEntity)
                .toList();
        return new ResponseEntity<>(platformNotifications, HttpStatus.OK);
    }

    @PostMapping("/create")
    public ResponseEntity<?> create(@RequestBody @Valid CreatePlatformNotificationRequest request) {
        PlatformNotificationEntity platformNotification = PlatformNotificationEntity.fromCreatePlatformNotificationRequest(request);
        String id = platformNotificationService.create(platformNotification);
        CreatePlatformNotificationResponse response = new CreatePlatformNotificationResponse(id);
        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    @PutMapping("/check/{notificationId}")
    public ResponseEntity<?> check(@PathVariable String notificationId) {
        platformNotificationService.check(notificationId);
        return SuccessResponse.create200("Platform notification is checked");
    }
}
