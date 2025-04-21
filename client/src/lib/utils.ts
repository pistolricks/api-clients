import {type ClassValue, clsx} from "clsx"
import {twMerge} from "tailwind-merge"

import image1 from "~/img/cta_group_1.webp";

import image2 from "~/img/cta_group_1.webp";
import image3 from "~/img/cta_group_1.webp";

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs))
}

export type Bases = "default" | "custom"
export type Variants = "default" | "secondary" | "outline" | "ghost" | "destructive" | "link" | null | undefined

export type Sizes = "default" | "sm" | "lg" | "icon" | null | undefined

export function handleInitials(string?: string) {
    if (!string) return;
    return string?.split(' ')?.map((sp) => sp[0])
}

export function handleSplit(string?: string) {
    if (!string) return;
    return string?.trim().split(' ')
}

export function handleCamelCase(arr?: string[]) {
    if (!arr) return;

    let newArr = [];

    for (let i = 0; i < arr?.length; i++) {
        let ns = arr[i]?.trim();
        let str = ns?.[0]?.toUpperCase() + ns?.substring(1);

        newArr.push(str);
    }

    return newArr.join('');
}



export function handleUserName(string?: string) {
    if(!string) return;
    let split = handleSplit(string);
    let camelCase = handleCamelCase(split);
    return "@" + camelCase;
}

export function classNames(...classes: (string | boolean | undefined)[]): string {
    return classes.filter(Boolean).join(' ');
}

export function throttle(mainFunction: any, delay: number | undefined) {
    let timerFlag: ReturnType<typeof setTimeout> | null = null; // Variable to keep track of the timer

    // Returning a throttled version
    return (...args: any) => {
        if (timerFlag === null) { // If there is no timer currently running
            mainFunction(...args as [any]); // Execute the main function
            timerFlag = setTimeout(() => { // Set a timer to clear the timerFlag after the specified delay
                timerFlag = null; // Clear the timerFlag to allow the main function to be executed again
            }, delay);
        }
    };
}


export function arrayDedupe<T>(array: T[], b: T, key: keyof T) {
    let items: T[] = array.filter((c: T) => c[key] !== b[key]);
    return [...items, b].reverse();

}

export function arrayRemove<T>(array: T[], b: T, key: keyof T) {
    let items: T[] = array.filter((c: T) => c[key] !== b[key]);
    return items;

}


export function capitalizeFirstLetter(str?: string) {
    if(!str) return;
    return str[0].toUpperCase() + str.slice(1);
}

export function lowerCaseString(str?: string) {
    if(!str) return;
    return str.toLowerCase();
}

export const renameKey = (oldKey: string, newKey: string, map: Map<string, any>) => {
    const old = map.get(oldKey);
    const rest = [...map].filter(r => r[0] !== oldKey);
    return new Map(
        [[newKey, old],
            ...rest]);
}


