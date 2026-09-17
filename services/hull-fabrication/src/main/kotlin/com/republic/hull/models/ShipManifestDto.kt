package com.republic.hull.models

import java.time.LocalDateTime
import java.util.UUID

/**
 * The clean data structure returned to the API consumers.
 * This prevents us from accidentally exposing database internals.
 */
data class ShipManifestDto(
    val manifestId: UUID,
    val className: String,
    val status: String,
    val fabricationTimestamp: LocalDateTime
)

/**
 * The incoming request payload when a new ship is ordered.
 */
data class ShipBlueprintRequest(
    val className: String,
    val hullAlloy: String,
    val deflectorShields: Boolean = true
)
