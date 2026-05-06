import { useCallback, useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import useEmblaCarousel from 'embla-carousel-react';

function PhotoCarousel({ photos, onPhotoClick }: { photos: string[]; onPhotoClick: (i: number) => void }) {
  const [emblaRef, emblaApi] = useEmblaCarousel({ loop: true });
  const [thumbRef, thumbApi] = useEmblaCarousel({ containScroll: 'keepSnaps', dragFree: true });
  const [selected, setSelected] = useState(0);

  const onThumbClick = useCallback((i: number) => {
    emblaApi?.scrollTo(i);
  }, [emblaApi]);

  useEffect(() => {
    if (!emblaApi) return;
    const onSelect = () => {
      const i = emblaApi.selectedScrollSnap();
      setSelected(i);
      thumbApi?.scrollTo(i);
    };
    emblaApi.on('select', onSelect);
    return () => { emblaApi.off('select', onSelect); };
  }, [emblaApi, thumbApi]);

  return (
    <div>
      <div className="relative">
        <div className="overflow-hidden rounded-2xl" ref={emblaRef}>
          <div className="flex">
            {photos.map((src, i) => (
              <div key={i} className="flex-none w-full">
                <img
                  src={src}
                  alt={`photo ${i + 1}`}
                  className="w-full aspect-video object-cover cursor-pointer"
                  onClick={() => onPhotoClick(i)}
                />
              </div>
            ))}
          </div>
        </div>
        <button
          onClick={() => emblaApi?.scrollPrev()}
          className="absolute left-3 top-1/2 -translate-y-1/2 w-12 h-12 rounded-full bg-black/40 hover:bg-brand/60 flex items-center justify-center transition"
        >
          <img src="/static/arrow-left.svg" alt="prev" className="w-6 h-6" />
        </button>
        <button
          onClick={() => emblaApi?.scrollNext()}
          className="absolute right-3 top-1/2 -translate-y-1/2 w-12 h-12 rounded-full bg-black/40 hover:bg-brand/60 flex items-center justify-center transition"
        >
          <img src="/static/arrow-right.svg" alt="next" className="w-6 h-6" />
        </button>
      </div>
      <div className="mt-3" ref={thumbRef}>
        <div className="flex gap-2">
          {photos.map((src, i) => (
            <button
              key={i}
              onClick={() => onThumbClick(i)}
              className={`flex-1 rounded-lg overflow-hidden border-2 transition ${
                selected === i ? 'border-brand' : 'border-transparent opacity-50 hover:opacity-80'
              }`}
            >
              <img src={src} alt={`thumb ${i + 1}`} className="w-full aspect-video object-cover" />
            </button>
          ))}
        </div>
      </div>
    </div>
  );
}
import type { Vehicle } from '../types/vehicle';
import { fetchVehicle } from '../api/vehicles';
import Lightbox from '../components/Lightbox';

export default function VehicleDetailsPage() {
  const { id } = useParams<{ id: string }>();
  const { t } = useTranslation();
  const [v, setV] = useState<Vehicle | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [lbOpen, setLbOpen] = useState(false);
  const [lbStart, setLbStart] = useState(0);

  useEffect(() => {
    if (!id) return;
    fetchVehicle(Number(id))
      .then(setV)
      .catch(e => setError(String(e)));
  }, [id]);

  if (error) return <div className="text-center py-20 text-red-400">{error}</div>;
  if (!v)    return <div className="text-center py-20 text-white/50">Loading...</div>;

  const fixPath = (p: string) => p.replace(/^\.\//, '/');
  const allPhotos = [v.main_photo, ...(v.photos || [])].map(fixPath);

  return (
    <section className="py-8">
      <div className="text-center mt-4 mb-8">
        <p className="section-title">{v.brand} {v.model}</p>
      </div>

      <div className="grid grid-cols-1 xl:grid-cols-3 gap-6">
        <div className="xl:col-span-2 bg-bg-dark rounded-2xl p-6">
          <PhotoCarousel
            photos={allPhotos}
            onPhotoClick={(i) => { setLbStart(i); setLbOpen(true); }}
          />
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

      <Lightbox
        isOpen={lbOpen}
        photos={allPhotos}
        initialIndex={lbStart}
        onClose={() => setLbOpen(false)}
      />
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
