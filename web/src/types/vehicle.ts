export interface Vehicle {
    id: number;
    brand: string;
    model: string;
    engine: string;
    transmission: string;
    hp_amount: number;
    fuel_type: string;
    year_of_prod: number;
    mileage: number;
    description: string; 
    main_photo: string;
    photos: string[];
}

export type VehicleInput = Omit<Vehicle, 'id' | 'main_photo' | 'photos'>;