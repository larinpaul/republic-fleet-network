package com.republic.hull.services

import com.republic.hull.models.HullStatus
import com.republic.hull.models.ShipBlueprintRequest
import com.republic.hull.models.ShipManifestDto
import com.republic.hull.models.StarshipHull
import com.republic.hull.repositories.StarshipHullRepository
import org.springframework.stereotype.Service
import org.springframework.transaction.annotation.Transactional
import java.util.UUID

/**
 * The business logic layer. 
 * Handles the actual "fabrication" of the ship data.
 */
/**
 * The business logic layer.
 * Hndles the actual "fabrication" of the ship data.
 * This Service is the foreman of the shipyard. 
 *  It takes the raw blueprints from the outside world,
 *  translates them into physical database entities,
 *  instructs the Repository to save them,
 *   and then formats the result back into a clea nmanifest for the pilot.
 */ 
@Service
class HullFabricationService(
    private val hullRepository: StarshipHullRepository
) {

    /**
     * Initiates the fabrication of a new starship hull.
     * The @Transactional annotation ensures that if anything fails 
     * during this process, the database rolls back (no partial ships!).
     */
    @Transactional
    fun fabricateShip(blueprint: ShipBlueprintRequest): ShipManifestDto {
        // 1. Translate the blueprint into a physical database entity
        val newHull = StarshipHull(
            className = blueprint.className,
            hullAlloy = blueprint.hullAlloy,
            deflectorShields = blueprint.deflectorShields,
            status = HullStatus.FABRICATING
        )

        // 2. Save to the PostgreSQL vault
        val savedHull = hullRepository.save(newHull)

        // 3. Return the clean DTO to the outside world
        return mapToDto(savedHull)
    }

    /**
     * Retrieves the status of a specific ship by its manifest ID.
     * Throws an exception if the ship doesn't exist (handled by the Controller).
     */
    fun getShipStatus(manifestId: UUID): ShipManifestDto {
        val hull = hullRepository.findById(manifestId)
            .orElseThrow { NoSuchElementException("Ship manifest $manifestId not found in Republic Archives.") }
        
        return mapToDto(hull)
    }

    /**
     * Internal helper to map the heavy Database Entity to the lightweight API DTO.
     */
    private fun mapToDto(hull: StarshipHull): ShipManifestDto {
        return ShipManifestDto(
            manifestId = hull.id,
            className = hull.className,
            status = hull.status.name,
            fabricationTimestamp = hull.fabricationTimestamp
        )
    }
}
