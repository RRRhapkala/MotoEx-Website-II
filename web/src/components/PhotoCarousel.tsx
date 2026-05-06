import { useCallback, useEffect, useState } from 'react';
import useEmblaCarousel from 'embla-carousel-react';

interface Props {
  photos: string[];
  onPhotoClick: (i: number) => void;
}

export default function PhotoCarousel({ photos, onPhotoClick }: Props) {
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
