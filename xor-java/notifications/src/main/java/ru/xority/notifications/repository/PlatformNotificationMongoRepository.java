package ru.xority.notifications.repository;

import java.util.List;
import java.util.Optional;

import lombok.RequiredArgsConstructor;
import org.bson.types.ObjectId;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.stereotype.Repository;

import ru.xority.notifications.entity.PlatformNotificationEntity;
import ru.xority.notifications.entity.PlatformNotificationFilter;
import ru.xority.notifications.entity.PlatformNotificationMongoEntity;
import ru.xority.utils.DataFilterX;

/**
 * @author foxleren
 */
@Repository
@RequiredArgsConstructor
public class PlatformNotificationMongoRepository implements PlatformNotificationRepository {
    private static final String PLATFORM_NOTIFICATION_COLLECTION = "platform_notification";

    private final MongoTemplate mongoTemplate;

    @Override
    public List<PlatformNotificationEntity> list(PlatformNotificationFilter filter) {
        Query query = new Query();

        filter.getId().ifPresent(v -> query.addCriteria(Criteria.where(PlatformNotificationMongoEntity.ID_FIELD).is(v)));
        filter.getReceiverUuid().ifPresent(v -> query.addCriteria(Criteria.where(PlatformNotificationMongoEntity.RECEIVER_UUID_FIELD).is(v)));
        filter.getSenderUuid().ifPresent(v -> query.addCriteria(Criteria.where(PlatformNotificationMongoEntity.SENDER_UUID_FIELD).is(v)));

        return mongoTemplate.find(query, PlatformNotificationMongoEntity.class, PLATFORM_NOTIFICATION_COLLECTION)
                .stream()
                .map(PlatformNotificationEntity::fromPlatformNotificationMongoEntity)
                .toList();
    }

    @Override
    public Optional<PlatformNotificationEntity> get(PlatformNotificationFilter filter) {
        List<PlatformNotificationEntity> platformNotifications = this.list(filter);
        return DataFilterX.singleO(platformNotifications);
    }

    @Override
    public String create(PlatformNotificationEntity platformNotification) {
        PlatformNotificationMongoEntity platformNotificationMongo = PlatformNotificationMongoEntity
                .fromPlatformNotificationEntity(platformNotification);
        return mongoTemplate.insert(platformNotificationMongo, PLATFORM_NOTIFICATION_COLLECTION)
                .getId().toHexString();
    }

    @Override
    public void update(PlatformNotificationEntity platformNotification) {
        PlatformNotificationMongoEntity platformNotificationMongo = PlatformNotificationMongoEntity
                .fromPlatformNotificationEntity(platformNotification);

        Query query = new Query();

        ObjectId id = new ObjectId(platformNotificationMongo.getId().toHexString());
        query.addCriteria(Criteria.where(PlatformNotificationMongoEntity.ID_FIELD).is(id));

        mongoTemplate.replace(query, platformNotificationMongo, PLATFORM_NOTIFICATION_COLLECTION);
    }
}
