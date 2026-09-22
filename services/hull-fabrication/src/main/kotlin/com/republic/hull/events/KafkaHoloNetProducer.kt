package com.republic.hull.events

import org.slf4j.LoggerFactory
import org.springframework.kafka.core.KafkaTemplate
import org.springframework.stereotype.Component

/**
 * This is the actualcomms array.
 * It uses Spring's KafkaTemplate to push messages to a specific Kafka topic.
 * Notice how we use the manifestId as the Kafka "Key".
 * This ensures that all events for the exact same ship 
 * are routed to the same partition,
 * keeping the timeline of that specific ship pefectly in order

 * The Kafka Producer. 
 * Responsible for broadcasting events to the Republic HoloNet.
 */
@Component
class KafkaHoloNetProducer(
    private val kafkaTemplate: KafkaTemplate<String, Any>
) {
    private val logger = LoggerFactory.getLogger(javaClass)

    // The specific frequency (topic) we broadcast hull completions on
    companion object {
        const val HULL_COMPLETED_TOPIC = "republic.ship.hull.completed"
    }

    /**
     * Transmits the HullCompletedEvent to the Kafka cluster.
     */
    fun publishHullCompleted(event: HullCompletedEvent) {
        val key = event.manifestId.toString()
        
        logger.info("Broadcasting to HoloNet: Hull {} completed. Transmitting on topic '{}'", key, HULL_COMPLETED_TOPIC)

        // Send the message. We use the manifestId as the Kafka Key for partition routing.
        val future = kafkaTemplate.send(HULL_COMPLETED_TOPIC, key, event)

        // Async callback to log if the transmission succeeded or failed
        future.whenComplete { result, ex ->
            if (ex == null) {
                logger.info("Transmission successful. Offset: {}", result.recordMetadata.offset())
            } else {
                logger.error("HoloNet transmission FAILED for hull {}!", key, ex)
            }
        }
    }
}