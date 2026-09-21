package com.republic.hull.events

import java.time.LocalDateTime
import java.util.UUID

/**
 * Before we broadcast, we need to format the message.
 * This is the exact JSON structure that wil travel across the Kafka topic.
 * We keep it lightweight - only what the other services actually need.
 * The payload broadcasted to the HoloNet (Kafka) when a hull is finished.
 * This is serialized into JSON and sent over the wire.
 */
data class HullCompletedEvent(
    val manifestId: UUID,
    val className: String,
    val hullAlloy: String,
    val completedAt: LocalDateTime
)
