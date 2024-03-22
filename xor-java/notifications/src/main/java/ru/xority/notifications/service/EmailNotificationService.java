package ru.xority.notifications.service;

import java.nio.charset.StandardCharsets;
import java.time.LocalDateTime;
import java.util.List;
import java.util.UUID;

import jakarta.mail.internet.MimeMessage;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.javamail.MimeMessageHelper;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;
import org.thymeleaf.context.Context;
import org.thymeleaf.spring6.SpringTemplateEngine;

import ru.xority.notifications.entity.EmailNotificationEntity;
import ru.xority.notifications.entity.EmailNotificationFilter;
import ru.xority.notifications.repository.EmailNotificationRepository;

/**
 * @author foxleren
 */
@Service
@RequiredArgsConstructor
public class EmailNotificationService {
    private static final String EMAIL_TEMPLATE_PATH = "email_template.html";

    private final JavaMailSender emailSender;
    private final SpringTemplateEngine templateEngine;
    private final EmailNotificationRepository emailNotificationRepository;

    @Value("${email-notification.sender-email}")
    private String senderEmail;

    public List<EmailNotificationEntity> list(EmailNotificationFilter filter) {
        return emailNotificationRepository.list(filter);
    }

    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public UUID create(
            UUID senderUuid,
            String receiverEmail,
            String subject,
            String body) {
        UUID notificationUuid = UUID.randomUUID();
        EmailNotificationEntity entity = new EmailNotificationEntity(
                notificationUuid,
                senderUuid,
                senderEmail,
                receiverEmail,
                subject,
                body,
                LocalDateTime.now()
        );
        emailNotificationRepository.create(entity);
        sendEmail(receiverEmail, subject, body);
        return notificationUuid;
    }

    private void sendEmail(String receiverEmail,
                           String subject,
                           String body) {
        Context context = new Context();
        context.setVariable("EMAIL_TITLE", subject);
        context.setVariable("EMAIL_CONTENT", body);
        String template = templateEngine.process(EMAIL_TEMPLATE_PATH, context);

        try {
            MimeMessage message = emailSender.createMimeMessage();
            MimeMessageHelper mimeMessageHelper = new MimeMessageHelper(
                    message,
                    MimeMessageHelper.MULTIPART_MODE_MIXED_RELATED,
                    StandardCharsets.UTF_8.name()
            );

            mimeMessageHelper.setTo(receiverEmail);
            mimeMessageHelper.setSubject(subject);
            mimeMessageHelper.setFrom(senderEmail);
            mimeMessageHelper.setText(template, true);
            emailSender.send(message);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}
