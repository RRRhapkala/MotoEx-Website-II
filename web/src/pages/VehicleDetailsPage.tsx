import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import type { Vehicle } from '../types/vehicle';
import { fetchVehicle } from '../api/vehicles';

export default function VehicleDetailsPage() {
  const { id } = useParams<{ id: string }>();
  const { t } = useTranslation();
  const [v, setV] = useState<Vehicle | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!id) return;
    fetchVehicle(Number(id))
      .then(setV)
      .catch(e => setError(String(e)));
  }, [id]);

  if (error)  return <div className="text-center py-20 text-red-400">{error}</div>;
  if (!v)     return <div className="text-center py-20 text-white/50">Loading...</div>;

  const fixPath = (p: string) => p.replace(/^\.\//, '/');
  const allPhotos = [v.main_photo, ...(v.photos || [])].map(fixPath);

  return (
    <section className="py-8">
      <div className="text-center mt-4 mb-8">
        <p className="section-title">{v.brand} {v.model}</p>
      </div>
      <div className="grid grid-cols-1 xl:grid-cols-3 gap-6">
        <div className="xl:col-span-2 bg-bg-dark rounded-2xl p-6">
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
            {allPhotos.map((src, i) => (
              <img
                key={i}
                src={src}
                alt={`${v.brand} ${v.model} photo ${i + 1}`}
                className="w-full aspect-[4/3] object-cover rounded-xl border-2 border-brand/20"
                loading="lazy"
              />
            ))}
          </div>
        </div>
        <div className="bg-bg-dark rounded-2xl p-6 border border-brand/10">
          <InfoRow label={t('engine')}     value={v.engine} />
          <InfoRow label={t('fuel_type')}  value={v.fuel_type} />
          <InfoRow label={t('gearbox')}    value={v.transmission} />
          <InfoRow label={t('horsepower')} value={`${v.hp_amount} HP`} />
          <InfoRow label={t('year')}       value={String(v.year_of_prod)} />
          <InfoRow label={t('mileage')}    value={`${v.mileage.toLocaleString()} km`} />
          <h3 className="text-brand font-semibold text-xl mt-6 mb-3">{t('description')}</h3>
          <p className="text-white/85 leading-relaxed text-justify">{v.description}</p>
        </div>
      </div>
    </section>
  );
}

function InfoRow({ label, value }: { label: string; value: string }) {
  return (
    <div className="flex items-center py-3 border-b border-white/10 last:border-0">
      <span className="flex-1 text-white/70">{label}</span>
      <span className="text-white">{value}</span>
    </div>
  );
}