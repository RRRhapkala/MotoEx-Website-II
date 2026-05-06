import type { Vehicle, VehicleInput } from '../types/vehicle';
import { apiGet, apiPost, apiPut, apiDelete } from './client';

export const fetchVehicles = ()      => apiGet<Vehicle[]>('/cars');
export const fetchVehicle  = (id: number) => apiGet<Vehicle>(`/cars/${id}`);
export const createVehicle = (v: VehicleInput, token: string) =>
  apiPost<Vehicle>('/admin', v, token);
export const updateVehicle = (id: number, v: Vehicle, token: string) =>
  apiPut<Vehicle>(`/admin/${id}`, v, token);
export const deleteVehicle = (id: number, token: string) =>
  apiDelete<{ message: string }>(`/admin/${id}`, token);