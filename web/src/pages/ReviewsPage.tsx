import { useTranslation } from 'react-i18next';

interface Review {
  author: string;
  date: string;
  rating: number;
  title: string;
  body: string;
  tag: string;
}

const REVIEWS: Review[] = [
  {
    author: 'Marcin Kowalski', date: 'January 2025', rating: 5,
    title: 'Perfect import from Germany — exceeded expectations',
    body: 'Ordered a BMW 5 Series from Germany. The whole process was incredibly smooth — from selecting the car to delivery at my door. MotoEx handled every single document, customs, and registration without any issues. I was kept informed at every step. The car arrived in perfect condition, exactly as described. Absolutely recommend.',
    tag: 'BMW · Germany',
  },
  {
    author: 'Anna Wiśniewska', date: 'November 2024', rating: 5,
    title: 'Fast, transparent, and professional service',
    body: 'I was a bit hesitant at first about importing a car, but MotoEx made it feel effortless. They found the exact Audi A6 I was looking for in the Netherlands within a week. Pricing was fully transparent — no hidden fees. Communication was excellent throughout the whole process. Will definitely use them again.',
    tag: 'Audi · Netherlands',
  },
  {
    author: 'Piotr Nowak', date: 'October 2024', rating: 4,
    title: 'Great experience, minor delay but resolved quickly',
    body: 'Overall a really positive experience with MotoEx. My Mercedes C-Class from Belgium arrived about 5 days later than the initial estimate due to transport issues — but the team communicated proactively and sorted it out fast. The car itself was in outstanding condition. Minus one star for the delay, but I\'d use them again.',
    tag: 'Mercedes · Belgium',
  },
  {
    author: 'Tomasz Jabłoński', date: 'August 2024', rating: 5,
    title: 'Second car imported through MotoEx — loyal customer',
    body: 'This is my second import with MotoEx and they delivered once again. This time a Porsche Cayenne from Austria. They negotiated a better price than I could find locally and managed the entire logistics chain. Everything from inspection to registration was handled professionally. These guys are the real deal.',
    tag: 'Porsche · Austria',
  },
  {
    author: 'Karolina Dąbrowska', date: 'June 2024', rating: 5,
    title: 'Saved me a lot of money compared to local dealerships',
    body: 'I compared the total import cost (car + MotoEx fees + all duties) with local dealership prices for the same Volkswagen Tiguan. Saved over 15,000 PLN. The service was impeccable — detailed inspection report with photos before purchase, transparent cost breakdown, and delivery on time. Highly satisfied.',
    tag: 'Volkswagen · Germany',
  },
  {
    author: 'Rafał Wróblewski', date: 'March 2024', rating: 5,
    title: 'Exceptional attention to detail throughout the process',
    body: 'From the very first consultation, I could tell MotoEx operates on a different level. They helped me find a rare BMW M4 Competition in France, verified its full history, and had it transported with premium enclosed transport. Not a single scratch. The whole journey took 3 weeks. Worth every penny.',
    tag: 'BMW M4 · France',
  },
];

function Stars({ rating }: { rating: number }) {
  return (
    <div className="flex gap-1">
      {[1,2,3,4,5].map(i => (
        <span key={i} style={{ color: i <= rating ? '#ff6600' : 'rgba(255,255,255,0.2)', fontSize: 18 }}>★</span>
      ))}
    </div>
  );
}

function ReviewCard({ r }: { r: Review }) {
  return (
    <div className="flex flex-col gap-3 rounded-2xl p-7 border border-brand/10 hover:-translate-y-1 hover:border-brand/35 hover:shadow-2xl transition"
      style={{ background: 'rgba(45,45,45,0.9)' }}>
      <div className="flex items-center gap-4">
        <div className="w-12 h-12 rounded-full flex items-center justify-center flex-shrink-0"
          style={{ background: 'linear-gradient(135deg, #ff6600, #ff8533)' }}>
          <span className="text-white text-xl">👤</span>
        </div>
        <div className="flex-1 min-w-0">
          <div className="text-white text-base truncate">{r.author}</div>
          <div className="text-white/40 text-xs tracking-wide mt-0.5">{r.date}</div>
        </div>
      </div>
      <Stars rating={r.rating} />
      <hr className="border-none h-px" style={{ background: 'rgba(255,102,0,0.12)' }} />
      <div className="text-white font-semibold text-base leading-snug">{r.title}</div>
      <p className="text-white/75 text-sm leading-[1.8] text-justify flex-1">{r.body}</p>
      <span className="self-start text-xs tracking-widest px-3 py-1 rounded-full border border-brand/20"
        style={{ background: 'rgba(255,102,0,0.12)', color: '#ff8533' }}>
        {r.tag}
      </span>
    </div>
  );
}

export default function ReviewsPage() {
  const { t } = useTranslation();

  return (
    <section className="py-8">
      <div className="text-center mt-4 mb-2">
        <p className="section-title">{t('customer_reviews')}</p>
      </div>
      <div className="text-center mb-10 text-white/50 tracking-widest text-sm uppercase">
        {t('customer_reviews_text')}
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-10">
        {REVIEWS.map((r, i) => <ReviewCard key={i} r={r} />)}
      </div>
    </section>
  );
}
