import { useTranslation } from 'react-i18next'

export default function LangSwitcher() {
    const { i18n } = useTranslation();
    const change = (l: string) => {
        i18n.changeLanguage(l);
        localStorage.setItem('lang', l)
    };
    const cls = (l: string) =>
    `px-2 py-1 rounded ${i18n.language === l ? 'text-brand font-semibold' : 'text-white/50 hover:text-white'}`;
  return (
    <div className="flex gap-2 ml-6">
      <button className={cls('pl')} onClick={() => change('pl')}>PL</button>
      <span className="text-white/20">|</span>
      <button className={cls('ru')} onClick={() => change('ru')}>RU</button>
      <span className="text-white/20">|</span>
      <button className={cls('en')} onClick={() => change('en')}>EN</button>
    </div>
  );
}