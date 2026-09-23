package com.republic.hull.services

import com.republic.hull.events.HullCompletedEvent
import com.republic.hull.events.KafkaHoloNetProducer
import com.republic.hull.models.HullStatus
import com.republic.hull.models.ShipBlueprintRequest
import com.republic.hull.models.ShipManifestDto
import com.republic.hull.models.StarshipHull
import com.republic.hull.repositories.StarshipHullRepository
import org.springframework.stereotype.Service
import org.springframework.transaction.annotation.Transactional
import java.time.LocalDateTime
import java.util.UUID

@Service
class HullFabricationService(
    private val hullRepository: StarshipHullRepository,
    private val holoNetProducer: KafkaHoloNetProducer // <-- Injected the HoloNet transmitter
) {

    @Transactional
    fun fabricateShip(blueprint: ShipBlueprintRequest): ShipManifestDto {
        // 1. Create the physical entity
        val newHull = StarshipHull(
            className = blueprint.className,
            hullAlloy = blueprint.hullAlloy,
            deflectorShields = blueprint.deflectorShields,
            status = HullStatus.HULL_COMPLETE_AWAITING_HYPERDRIVE // Updated status!
        )

        // 2. Save to the PostgreSQL vault
        val savedHull = hullRepository.save(newHull)

        // 3. BROADCAST TO THE HOLONET! (Kafka)
        val event = HullCompletedEvent(
            manifestId = savedHull.id,
            className = savedHull.className,
            hullAlloy = savedHull.hullAlloy,
            completedAt = LocalDateTime.now()
        )
        holoNetProducer.publishHullCompleted(event)

        // 4. Return the clean DTO to the API caller
        return mapToDto(savedHull)
    }

    fun getShipStatus(manifestId: UUID): ShipManifestDto {
        val hull = hullRepository.findById(manifestId)
            .orElseThrow { NoSuchElementException("Ship manifest $manifestId not found in Republic Archives.") }
        
        return mapToDto(hull)
    }

    private fun mapToDto(hull: StarshipHull): ShipManifestDto {
        return ShipManifestDto(
            manifestId = hull.id,
            className = hull.className,
            status = hull.status.name,
            fabricationTimestamp = hull.fabricationTimestamp
        )
    }
}
