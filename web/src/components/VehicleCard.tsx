import { Link } from 'react-router-dom';
import type { Vehicle } from '../types/vehicle';

export default function VehicleCard({ vehicle }: { vehicle: Vehicle }) {
  const photo = vehicle.main_photo.replace(/^\.\//, '/');
  return (
    <Link
      to={`/about/${vehicle.id}`}
      className="block bg-bg-dark rounded-2xl p-5 border border-brand/10 shadow-lg hover:-translate-y-1 transition"
    >
      <img
        src={photo}
        alt={`${vehicle.brand} ${vehicle.model}`}
        className="w-full aspect-[14/9] object-cover rounded-xl"
        loading="lazy"
      />
      <p className="text-center text-lg pt-4">{vehicle.brand} {vehicle.model}</p>
    </Link>
  );
}