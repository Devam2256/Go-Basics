# Go-Basics
this is a go basics repository

https://excalidraw.com/#json=kSZ7Hz2LYLqVxafvHzWiI,6lNzp8Ap-jYXSx7DGTuIRQ


// "use client";
// import { motion } from "motion/react";

// type IsometricProps = {
//     className?: string;
//     variant: "top" | "left" | "right";
//     isActive?: boolean;
// };

// export const Isometric = ({ className, variant, isActive }: IsometricProps) => {
//   const TRANSITION = {
//     type: "spring" as const,
//     stiffness: 300,
//     damping: 30,
//   };

//   const YVariants = {
//     animate: {
//         translateY: -20,
//     },
//     initial: {
//         translateY: 0,
//     },
//   };

//   const XVariants = {
//     animate: {
//         translateX: 20,
//     },
//     initial: {
//         translateX: 0,
//     },
//   };

//   const NegativeXVariants = {
//     animate: {
//         translateX: -20,
//     },
//     initial: {
//         translateX: 0,
//     },
//   };

//   const noOptVariants = {
//     animate: {
//         translateX: 0,
//         translateY: 0,
//     },
//     initial: {
//         translateX: 0,
//         translateY: 0,
//     },
//   };

//   const getVariants = (face: "top" | "left" | "right") => {
//     if(face != variant) return noOptVariants;

//     switch(face) {
//         case "top":
//             return YVariants;
//         case "left":
//             return NegativeXVariants;
//         case "right":
//             return XVariants;
//         default:
//             return noOptVariants;   
//     }
//   }

//   const getStrokeColor = (face: "top" | "left" | "right") => {
//     if(face != variant) return "var(--color-neutral-200)";
//     if(!isActive) return "var(--color-neutral-200)";

//     switch(face) {
//         case "top":
//             return "var(--color-blue-500)";
//         case "left":
//             return "var(--color-blue-500)";
//         case "right":
//             return "var(--color-blue-500)";
//         default:
//             return "var(--color-neutral-200)";
//     }
//  };

//   return (
//     <motion.div whileHover="animate" initial="initial">
//         <motion.svg viewBox="0 0 200 200" fill="none" xmlns="http://www.w3.org/2000/svg" className="size-60 text-neutral-200 dark:text-neutral-800">
//             <motion.path 
//                 variants={getVariants("top")}
//                 transition={TRANSITION}
//                 d="M100 40 Q108 40 155 68 Q162 72 155 76 Q108 104 100 104 Q92 104 45 76 Q38 72 45 68 Q92 40 100 40 Z"
//                 className="fill-neutral-50 dark:fill-neutral-900"
//                 stroke={getStrokeColor("top")}
//                 strokeWidth="1.5"
//             />

//             <motion.path
//                 variants={getVariants("top")}
//                 transition={TRANSITION}
//                 d="M100 52 Q105 52 132 68 Q138 72 132 76 Q105 92 100 92 Q95 92 68 76 Q62 72 68 68 Q95 52 100 52 Z"
//                 className="fill-white dark:fill-neutral-950"
//                 stroke={getStrokeColor("top")}
//                 strokeWidth="1"
//             />

//             <motion.path
//                 variants={getVariants("left")}
//                 transition={TRANSITION}
//                 d="M45 76 L100 104 L100 164 Q100 170 92 166 L45 140 Q38 136 38 128 L38 80 Q38 72 45 76 Z"
//                 className="fill-neutral-100 dark:fill-neutral-950"
//                 stroke={getStrokeColor("left")}
//                 strokeWidth="1.5"
//             />
            
//             <motion.path
//                 variants={getVariants("right")}
//                 transition={TRANSITION}
//                 d="M155 76 L100 104 L100 164 Q100 170 108 166 L155 140 Q162 136 162 128 L162 80 Q162 72 155 76 Z"
//                 className="fill-neutral-50 dark:fill-neutral-900"
//                 stroke={getStrokeColor("right")}
//                 strokeWidth="1.5"
//             />

//             <motion.path
//                 variants={getVariants("left")}
//                 transition={TRANSITION}
//                 d="M55 86 L55 145"
//                 stroke={getStrokeColor("left")}
//                 strokeWidth="1"
//                 strokeLinecap="round"
//                 strokeDasharray="3 3"
//             />
            
//             <motion.path
//                 variants={getVariants("left")}
//                 transition={TRANSITION}
//                 d="M70 95 L70 155"
//                 stroke={getStrokeColor("left")}
//                 strokeWidth="1"
//                 strokeLinecap="round"
//                 strokeDasharray="3 3"
//             />

//             <motion.path
//                 variants={getVariants("left")}
//                 transition={TRANSITION}
//                 d="M85 104 L85 162"
//                 stroke={getStrokeColor("left")}
//                 strokeWidth="1"
//                 strokeLinecap="round"
//                 strokeDasharray="3 3"
//             />
            
//             <motion.path
//                 variants={getVariants("right")}
//                 transition={TRANSITION}
//                 d="M115 104 L115 162"
//                 stroke={getStrokeColor("right")}
//                 strokeWidth="1"
//                 strokeLinecap="round"
//                 strokeDasharray="3 3"
//             />

//             <motion.path
//                 variants={getVariants("right")}
//                 transition={TRANSITION}
//                 d="M130 95 L130 155"
//                 stroke={getStrokeColor("right")}
//                 strokeWidth="1"
//                 strokeLinecap="round"
//                 strokeDasharray="3 3"
//             />

//             <motion.path
//                 variants={getVariants("right")}
//                 transition={TRANSITION}
//                 d="M145 86 L145 145"
//                 stroke={getStrokeColor("right")}
//                 strokeWidth="1"
//                 strokeLinecap="round"
//                 strokeDasharray="3 3"
//             />
//         </motion.svg>
//     </motion.div>
//   );
// }


// ---------------------------------------------------------------------------------------------------------------------------------------------------

// "use client";
// import { motion, type Transition } from "motion/react";

// type CardStackProps = {
//   className?: string;
// };

// export const IsometricCardStack = ({ className }: CardStackProps) => {
//   const TRANSITION: Transition = {
//     duration: 0.5,
//     ease: [0.25, 0.8, 0.25, 1], // smooth ease (no spring)
//   };

//   /**
//    * Key idea:
//    * - All cards aligned diagonally (same angle)
//    * - Offset increases with depth
//    * - Hover = expand outward along same diagonal vector
//    * - Scale + opacity reinforce "coming out"
//    */

//   const bottom = {
//     initial: {
//       translateX: -16,
//       translateY: 12,
//       scale: 0.92,
//       rotateZ: -3,
//       rotateX: 12,
//       opacity: 0.7,
//     },
//     animate: {
//       translateX: -28,
//       translateY: 22,
//       scale: 0.96,
//       rotateZ: -3,
//       rotateX: 14,
//       opacity: 0.9,
//     },
//   };

//   const middle = {
//     initial: {
//       translateX: -8,
//       translateY: 6,
//       scale: 0.96,
//       rotateZ: -3,
//       rotateX: 10,
//       opacity: 0.85,
//     },
//     animate: {
//       translateX: -16,
//       translateY: 12,
//       scale: 1.0,
//       rotateZ: -3,
//       rotateX: 10,
//       opacity: 1,
//     },
//   };

//   const top = {
//     initial: {
//       translateX: 0,
//       translateY: 0,
//       scale: 1,
//       rotateZ: -3,
//       rotateX: 10,
//       opacity: 1,
//     },
//     animate: {
//       translateX: 8,
//       translateY: -6,
//       scale: 1.05,
//       rotateZ: -3,
//       rotateX: 10,
//       opacity: 1,
//     },
//   };

//   return (
//     <motion.div
//       whileHover="animate"
//       initial="initial"
//       className={className}
//       style={{ perspective: 1000 }}
//     >
//       <svg
//         viewBox="0 0 240 200"
//         xmlns="http://www.w3.org/2000/svg"
//         className="size-60"
//       >
//         {/* Bottom Card */}
//         <motion.rect
//           variants={bottom}
//           transition={TRANSITION}
//           x="60"
//           y="70"
//           width="130"
//           height="80"
//           rx="18"
//           className="fill-neutral-100 dark:fill-neutral-900"
//           stroke="var(--color-neutral-300)"
//           strokeWidth="1.2"
//         />

//         {/* Middle Card */}
//         <motion.rect
//           variants={middle}
//           transition={TRANSITION}
//           x="60"
//           y="70"
//           width="130"
//           height="80"
//           rx="18"
//           className="fill-neutral-50 dark:fill-neutral-800"
//           stroke="var(--color-neutral-300)"
//           strokeWidth="1.2"
//         />

//         {/* Top Card */}
//         <motion.g variants={top} transition={TRANSITION}>
//           <rect
//             x="60"
//             y="70"
//             width="130"
//             height="80"
//             rx="20"
//             className="fill-black"
//             stroke="var(--color-blue-500)"
//             strokeWidth="1.5"
//           />

//           {/* Subtle glow border */}
//           <rect
//             x="60"
//             y="70"
//             width="130"
//             height="80"
//             rx="20"
//             fill="none"
//             stroke="var(--color-blue-500)"
//             strokeWidth="1"
//             opacity="0.6"
//           />

//           {/* Chip */}
//           <rect
//             x="78"
//             y="92"
//             width="18"
//             height="14"
//             rx="3"
//             className="fill-neutral-600"
//           />

//           {/* Details */}
//           <line
//             x1="80"
//             y1="120"
//             x2="165"
//             y2="120"
//             stroke="var(--color-blue-500)"
//             strokeWidth="1.2"
//             strokeDasharray="5 4"
//             strokeLinecap="round"
//           />

//           <line
//             x1="80"
//             y1="132"
//             x2="140"
//             y2="132"
//             stroke="var(--color-blue-500)"
//             strokeWidth="1.2"
//             strokeDasharray="5 4"
//             strokeLinecap="round"
//           />
//         </motion.g>
//       </svg>
//     </motion.div>
//   );
// };

// ---------------------------------------------------------------------------------------------------------------------------------------------------

// "use client";
// import { motion, type Transition } from "motion/react";

// type ShieldProps = {
//   className?: string;
// };

// const SHIELD = "M100 22 L152 48 L152 98 Q152 145 100 178 Q48 145 48 98 L48 48 Z";
// const INNER  = "M100 36 L140 56 L140 97 Q140 133 100 162 Q60 133 60 97 L60 56 Z";
// const PULSE  = "M100 50 Q132 62 132 95 Q132 128 100 155 Q68 128 68 95 Q68 62 100 50 Z";

// const TRANSITION: Transition = {
//   duration: 0.55,
//   ease: [0.25, 0.8, 0.25, 1],
// };

// export const IsometricShield = ({ className }: ShieldProps) => {
//   return (
//     <motion.div
//       initial="idle"
//       whileHover="hover"
//       className={className}
//     >
//       <svg viewBox="0 0 200 200" className="size-60 overflow-visible">

//         {/* Back layer */}
//         <motion.g
//           variants={{
//             idle:  { translateX: -22, translateY: 16, opacity: 0.38 },
//             hover: { translateX: 0,   translateY: 0,  opacity: 0 },
//           }}
//           transition={TRANSITION}
//         >
//           <path
//             d={SHIELD}
//             className="fill-neutral-100 dark:fill-neutral-900"
//             stroke="var(--color-neutral-300)"
//             strokeWidth="1.2"
//           />
//         </motion.g>

//         {/* Mid layer */}
//         <motion.g
//           variants={{
//             idle:  { translateX: -11, translateY: 8, opacity: 0.62 },
//             hover: { translateX: 0,   translateY: 0, opacity: 0 },
//           }}
//           transition={TRANSITION}
//         >
//           <path
//             d={SHIELD}
//             className="fill-neutral-50 dark:fill-neutral-800"
//             stroke="var(--color-neutral-200)"
//             strokeWidth="1.3"
//           />
//           <path d={INNER} fill="none" stroke="var(--color-neutral-200)" strokeWidth="0.8" />
//         </motion.g>

//         {/* Front layer */}
//         <motion.g
//           variants={{
//             idle:  { scale: 1,     opacity: 1 },
//             hover: { scale: 1.03,  opacity: 1 },
//           }}
//           style={{ originX: "100px", originY: "100px" }}
//           transition={TRANSITION}
//         >
//           <path
//             d={SHIELD}
//             className="fill-white dark:fill-neutral-950"
//             stroke="var(--color-blue-500)"
//             strokeWidth="1.6"
//           />
//           <path d={INNER} fill="none" stroke="var(--color-blue-500)" strokeWidth="1" opacity={0.55} />

//           {/* Diamond core mark */}
//           <path
//             d="M100 76 L113 92 L100 118 L87 92 Z"
//             fill="none"
//             stroke="var(--color-blue-500)"
//             strokeWidth="1.3"
//             strokeLinejoin="round"
//           />
//           <line x1="100" y1="82" x2="100" y2="112" stroke="var(--color-blue-500)" strokeWidth="0.8" opacity={0.4} />
//           <line x1="89"  y1="97" x2="111" y2="97"  stroke="var(--color-blue-500)" strokeWidth="0.8" opacity={0.4} />

//           {/* Pulse ring */}
//           <motion.path
//             d={PULSE}
//             fill="none"
//             stroke="var(--color-blue-500)"
//             strokeWidth="1.4"
//             variants={{
//               idle:  { opacity: 0, scale: 0.92 },
//               hover: {
//                 opacity: [0, 0.7, 0],
//                 scale:   [0.92, 1.0, 1.15],
//                 transition: { duration: 0.65, ease: "easeOut", delay: 0.4 },
//               },
//             }}
//             style={{ originX: "100px", originY: "97px" }}
//           />
//         </motion.g>
//       </svg>
//     </motion.div>
//   );
// };

// ----------------------------------------------------------------------------------------------------------------------------

"use client";
import { motion, type Transition } from "motion/react";

type Props = {
  className?: string;
};

export const IsometricSearchStack = ({ className }: Props) => {
  const TRANSITION: Transition = {
    duration: 0.22,
    ease: [0.4, 0, 0.2, 1],
  };

  const back = {
    initial: {
      translateX: -6,
      translateY: -6,
      scale: 0.94,
      rotateZ: -10,
      opacity: 0.4,
    },
    animate: {
      translateX: 0,
      translateY: 0,
      scale: 0.96,
      rotateZ: -10,
      opacity: 0.6,
    },
  };

  const mid = {
    initial: {
      translateX: -3,
      translateY: -3,
      scale: 0.97,
      rotateZ: -10,
      opacity: 0.7,
    },
    animate: {
      translateX: 0,
      translateY: 0,
      scale: 0.99,
      rotateZ: -10,
      opacity: 0.85,
    },
  };

  const front = {
    initial: {
      translateX: 0,
      translateY: 0,
      scale: 1,
      rotateZ: -10,
      opacity: 1,
    },
    animate: {
      translateX: 0,
      translateY: 0,
      scale: 1.02,
      rotateZ: -10,
      opacity: 1,
    },
  };

  return (
    <motion.div
      initial="initial"
      whileHover="animate"
      className={className}
      style={{ perspective: 1000 }}
    >
      <svg
        viewBox="0 0 200 200"
        xmlns="http://www.w3.org/2000/svg"
        className="size-60"
      >
        {/* BACK */}
        <motion.g variants={back} transition={TRANSITION}>
          <circle
            cx="90"
            cy="90"
            r="32"
            className="fill-neutral-100 dark:fill-neutral-900"
            stroke="var(--color-neutral-300)"
            strokeWidth="2"
          />
          <line
            x1="112"
            y1="112"
            x2="140"
            y2="140"
            stroke="var(--color-neutral-300)"
            strokeWidth="4"
            strokeLinecap="round"
          />
        </motion.g>

        {/* MID */}
        <motion.g variants={mid} transition={TRANSITION}>
          <circle
            cx="90"
            cy="90"
            r="32"
            className="fill-neutral-50 dark:fill-neutral-800"
            stroke="var(--color-neutral-300)"
            strokeWidth="2"
          />
          <line
            x1="112"
            y1="112"
            x2="140"
            y2="140"
            stroke="var(--color-neutral-300)"
            strokeWidth="4"
            strokeLinecap="round"
          />
        </motion.g>

        {/* FRONT */}
        <motion.g variants={front} transition={TRANSITION}>
          <motion.circle
            cx="90"
            cy="90"
            r="32"
            className="fill-white dark:fill-neutral-950"
            stroke="var(--color-blue-500)"
            strokeWidth="2.5"
            whileHover={{
              fill: "var(--color-blue-500)", // 👈 becomes bg-sky-500 equivalent
            }}
            transition={TRANSITION}
          />

          <motion.line
            x1="112"
            y1="112"
            x2="140"
            y2="140"
            stroke="var(--color-blue-500)"
            strokeWidth="4"
            strokeLinecap="round"
            whileHover={{
              stroke: "#ffffff", // contrast on blue
            }}
            transition={TRANSITION}
          />

          <motion.circle
            cx="90"
            cy="90"
            r="18"
            fill="none"
            stroke="var(--color-blue-500)"
            strokeWidth="1"
            opacity="0.5"
            whileHover={{
              stroke: "#ffffff",
              opacity: 0.7,
            }}
            transition={TRANSITION}
          />
        </motion.g>
      </svg>
    </motion.div>
  );
};
