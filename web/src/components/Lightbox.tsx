import { useEffect, useState } from 'react';

interface Props {
  isOpen: boolean;
  photos: string[];
  initialIndex: number;
  onClose: () => void;
}

export default function Lightbox({ isOpen, photos, initialIndex, onClose }: Props) {
  const [idx, setIdx] = useState(initialIndex);

  useEffect(() => { setIdx(initialIndex); }, [initialIndex]);

  useEffect(() => {
    if (!isOpen) return;
    const handler = (e: KeyboardEvent) => {
      if (e.key === 'Escape')    onClose();
      if (e.key === 'ArrowLeft') setIdx(i => (i - 1 + photos.length) % photos.length);
      if (e.key === 'ArrowRight') setIdx(i => (i + 1) % photos.length);
    };
    document.addEventListener('keydown', handler);
    document.body.style.overflow = 'hidden';
    return () => {
      document.removeEventListener('keydown', handler);
      document.body.style.overflow = '';
    };
  }, [isOpen, photos.length, onClose]);

  if (!isOpen) return null;

  return (
    <div
      className="fixed inset-0 bg-black/90 z-[9999] flex items-center justify-center cursor-zoom-out"
      onClick={onClose}
    >
      <span
        className="fixed top-5 right-7 text-4xl text-white/75 cursor-pointer hover:text-brand z-[10000]"
        onClick={onClose}
      >
        ×
      </span>
      <button
        className="fixed left-5 top-1/2 -translate-y-1/2 w-16 h-16 rounded-full bg-white/10 hover:bg-brand/30 z-[10000] flex items-center justify-center"
        onClick={(e) => { e.stopPropagation(); setIdx(i => (i - 1 + photos.length) % photos.length); }}
      >
        <img src="/static/arrow-left.svg" alt="prev" className="w-7 h-7" />
      </button>
      <img
        src={photos[idx]}
        alt=""
        className="max-w-[95vw] max-h-[92vh] rounded-xl object-contain shadow-2xl"
        onClick={(e) => e.stopPropagation()}
      />
      <button
        className="fixed right-5 top-1/2 -translate-y-1/2 w-16 h-16 rounded-full bg-white/10 hover:bg-brand/30 z-[10000] flex items-center justify-center"
        onClick={(e) => { e.stopPropagation(); setIdx(i => (i + 1) % photos.length); }}
      >
        <img src="/static/arrow-right.svg" alt="next" className="w-7 h-7" />
      </button>
      <span className="fixed bottom-5 left-1/2 -translate-x-1/2 text-white/60 text-sm tracking-widest">
        {idx + 1} / {photos.length}
      </span>
    </div>
  );
}