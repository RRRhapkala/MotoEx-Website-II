import { Outlet, Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import LangSwitcher from '../i18n/LangSwitcher';

function Navbar() {
  const { t } = useTranslation();

  return (
    <nav className="sticky top-0 z-50 bg-bg-darker/95 backdrop-blur border-b border-white/5 px-6 py-3">
      <div className="container mx-auto flex items-center gap-8">
        <Link to="/">
          <img src="/static/motoex-logo-true.svg" alt="MotoEx" className="h-10" />
        </Link>
        <div className="flex items-center gap-6 flex-1">
          <Link to="/catalog" className="text-white/70 hover:text-brand transition tracking-widest text-sm uppercase">
            {t('nav_catalog')}
          </Link>
          <Link to="/reviews" className="text-white/70 hover:text-brand transition tracking-widest text-sm uppercase">
            {t('nav_reviews')}
          </Link>
        </div>
        <LangSwitcher />
      </div>
    </nav>
  );
}

function Footer() {
  return (
    <footer className="bg-bg-darker border-t border-white/5 mt-10">
      <div className="container mx-auto px-6 py-8 grid grid-cols-1 md:grid-cols-3 gap-8">
        <div>
          <h4 className="text-white font-semibold mb-3 tracking-widest text-sm uppercase">About Us</h4>
          <p className="text-white/50 text-sm leading-relaxed">
            MotoEx — European car import company. We deliver quality vehicles directly from Western Europe.
          </p>
        </div>
        <div>
          <h4 className="text-white font-semibold mb-3 tracking-widest text-sm uppercase">Contact Us</h4>
          <ul className="space-y-2 text-white/50 text-sm">
            <li>✉ example@gmail.com</li>
            <li>📞 +48 784 213 831</li>
            <li>🏠 ul. Przykładowa 1, Warszawa</li>
          </ul>
        </div>
        <div>
          <h4 className="text-white font-semibold mb-3 tracking-widest text-sm uppercase">Follow Us</h4>
          <div className="flex gap-4 text-white/50">
            <a href="https://www.facebook.com/motoex2010" className="hover:text-brand transition">Facebook</a>
            <a href="https://www.instagram.com/motoex2010" className="hover:text-brand transition">Instagram</a>
            <a href="https://www.tiktok.com/@motoex2010" className="hover:text-brand transition">TikTok</a>
          </div>
        </div>
      </div>
    </footer>
  );
}

export default function Layout() {
  return (
    <div className="min-h-screen flex flex-col font-outfit text-white">
      <Navbar />
      <main className="flex-1 container mx-auto px-4">
        <Outlet />
      </main>
      <Footer />
    </div>
  );
}