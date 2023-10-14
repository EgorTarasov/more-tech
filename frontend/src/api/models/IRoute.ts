export interface IRoute {
    type: string;
    metadata: {
        attribution: string;
        service: string;
        timestamp: number;
        query: {
            coordinates: number[][];
            profile: string;
            format: string;
        };
        engine: {
            version: string;
            build_date: string;
            graph_date: string;
        };
    };
    bbox: number[];
    features: {
        bbox: number[];
        type: string;
        properties: {
            transfers: number;
            fare: number;
            segments: {
                distance: number;
                duration: number;
                steps: {
                    distance: number;
                    duration: number;
                    type: number;
                    instruction: string;
                    name: string;
                    way_points: number[];
                }[];
            }[];
            summary: {
                distance: number;
                duration: number;
            };
            way_points: number[];
        };
        geometry: {
            coordinates: number[][];
            type: string;
        };
    }[];
    routes: {
        distance: number;
        duration: number;
        geometry: {
            coordinates: number[][];
            type: string;
        };
        legs: {
            distance: number;
            duration: number;
            steps: {
                distance: number;
                duration: number;
                type: number;
                instruction: string;
                name: string;
                way_points: number[];
            }[];
            summary: {
                distance: number;
                duration: number;
            };
            way_points: number[];
        }[];
        weight: number;
        weight_name: string;
    }[];
}
