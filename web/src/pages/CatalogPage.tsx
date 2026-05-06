import { useEffect, useState } from 'react';
import type { Vehicle } from '../types/vehicle';
import { fetchVehicles } from '../api/vehicles';
import { ApiError } from '../api/client';
import { useTranslation } from 'react-i18next';
import VehicleCard from '../components/VehicleCard';

export default function CatalogPage() {
  const { t } = useTranslation();
  const [vehicles, setVehicles] = useState<Vehicle[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchVehicles()
      .then(setVehicles)
      .catch((e: ApiError) => setError(e.message))
      .finally(() => setLoading(false));
  }, []);

  return (
    <section>
      <div className="text-center mt-8 mb-8">
        <p className="section-title">{t('car_catalog')}</p>
      </div>
      {loading && <div className="text-center py-20 text-white/50">Loading...</div>}
      {error && <div className="text-center py-20 text-red-400">{error}</div>}
      {!loading && !error && vehicles.length === 0 && (
        <div className="text-center py-20 text-white/35">No vehicles in catalog yet</div>
      )}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mt-4 mb-10">
        {vehicles.map(v => <VehicleCard key={v.id} vehicle={v} />)}
      </div>
    </section>
  );
}