import { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import useEmblaCarousel from 'embla-carousel-react';
import Lightbox from '../components/Lightbox';

const CAROUSEL_IMAGES = [
  'https://images.unsplash.com/photo-1511919884226-fd3cad34687c?auto=format&fit=crop&q=80&w=1770',
  'https://images.unsplash.com/photo-1494976388531-d1058494cdd8?auto=format&fit=crop&q=80&w=1770',
  'https://images.unsplash.com/photo-1552519507-da3b142c6e3d?auto=format&fit=crop&q=80&w=1770',
  'https://images.unsplash.com/photo-1503376780353-7e6692767b70?auto=format&fit=crop&q=80&w=1770',
  'https://images.unsplash.com/photo-1542362567-b07e54358753?auto=format&fit=crop&q=80&w=1770',
  'https://images.unsplash.com/photo-1541443131876-44b03de101c5?auto=format&fit=crop&q=80&w=1770',
];

function Carousel() {
  const [emblaRef, emblaApi] = useEmblaCarousel({ loop: true });
  const [thumbRef, thumbApi] = useEmblaCarousel({ containScroll: 'keepSnaps', dragFree: true });
  const [selected, setSelected] = useState(0);
  const [lbOpen, setLbOpen] = useState(false);
  const [lbStart, setLbStart] = useState(0);

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
        <div className="overflow-hidden rounded-2xl" ref={emblaRef}
          style={{ border: '2px solid rgba(255,102,0,0.2)' }}>
          <div className="flex">
            {CAROUSEL_IMAGES.map((src, i) => (
              <div key={i} className="flex-none w-full">
                <img src={src} alt={`Car ${i + 1}`} loading="lazy" className="w-full aspect-video object-cover cursor-zoom-in"
                  onClick={() => { setLbStart(i); setLbOpen(true); }} />
              </div>
            ))}
          </div>
        </div>
        <button
          onClick={() => emblaApi?.scrollPrev()}
          className="absolute left-3 top-1/2 -translate-y-1/2 w-12 h-12 rounded-full bg-black/40 hover:bg-brand/60 flex items-center justify-center transition"
        >
          <img src="/static/arrow-left.svg" alt="prev" className="w-7 h-7" />
        </button>
        <button
          onClick={() => emblaApi?.scrollNext()}
          className="absolute right-3 top-1/2 -translate-y-1/2 w-12 h-12 rounded-full bg-black/40 hover:bg-brand/60 flex items-center justify-center transition"
        >
          <img src="/static/arrow-right.svg" alt="next" className="w-7 h-7" />
        </button>
      </div>

      <div className="mt-3" ref={thumbRef}>
        <div className="flex gap-2">
          {CAROUSEL_IMAGES.map((src, i) => (
            <button
              key={i}
              onClick={() => onThumbClick(i)}
              className={`flex-1 rounded-lg overflow-hidden border-2 transition ${
                selected === i ? 'border-brand' : 'border-transparent opacity-50 hover:opacity-80'
              }`}
            >
              <img src={src} alt={`Thumb ${i + 1}`} className="w-full aspect-video object-cover" />
            </button>
          ))}
        </div>
      </div>

      <Lightbox isOpen={lbOpen} photos={CAROUSEL_IMAGES} initialIndex={lbStart} onClose={() => setLbOpen(false)} />
    </div>
  );
}

export default function HomePage() {
  const { t } = useTranslation();

  return (
    <div className="py-4">
      <div className="text-center mt-4 mb-2">
        <p className="section-title">{t('how_works')}</p>
      </div>

      <div className="rounded-2xl shadow-lg my-8 py-1" style={{ background: '#3a3939' }}>
        <div className="rounded-2xl p-8 mx-5 my-6 border border-brand/10" style={{ background: 'rgba(45,45,45,0.9)' }}>
          <p className="text-white/90 leading-[1.8] text-justify" style={{ fontSize: 'clamp(16px, 2.5vh, 20px)' }}>
            {t('how_works_text')}
          </p>
        </div>
      </div>

      <div className="text-center mt-4 mb-6">
        <p className="section-title">{t('examples')}</p>
      </div>
      <div className="rounded-2xl shadow-lg my-4 py-1" style={{ background: '#3a3939' }}>
        <div className="rounded-2xl p-5 mx-5 my-6 border border-brand/10" style={{ background: 'rgba(45,45,45,0.9)' }}>
          <Carousel />
        </div>
      </div>
    </div>
  );
}
