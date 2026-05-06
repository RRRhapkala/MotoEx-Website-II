import { useEffect, useState } from 'react';
import type { Vehicle } from '../types/vehicle';
import { fetchVehicle } from '../api/vehicles';

export function useVehicle(id: string | undefined) {
  const [data, setData] = useState<Vehicle | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!id) return;
    let cancelled = false;
    setData(null);
    setError(null);
    setLoading(true);
    fetchVehicle(Number(id))
      .then(d  => { if (!cancelled) setData(d); })
      .catch(e => { if (!cancelled) setError(String(e)); })
      .finally(() => { if (!cancelled) setLoading(false); });
    return () => { cancelled = true; };
  }, [id]);

  return { data, loading, error };
}
