import { splitProps, getSlotCompoundVariant } from '../helpers.mjs';
import { createRecipe } from './create-recipe.mjs';

const tabsDefaultVariants = {
  "size": "md"
}
const tabsCompoundVariants = [
  {
    "size": "sm",
    "variant": "outline",
    "css": {
      "trigger": {
        "h": "9",
        "minW": "9",
        "textStyle": "sm",
        "px": "3.5"
      },
      "content": {
        "p": "3.5"
      }
    }
  },
  {
    "size": "md",
    "variant": "outline",
    "css": {
      "trigger": {
        "h": "10",
        "minW": "10",
        "textStyle": "sm",
        "px": "4"
      },
      "content": {
        "p": "4"
      }
    }
  },
  {
    "size": "lg",
    "variant": "outline",
    "css": {
      "trigger": {
        "h": "11",
        "minW": "11",
        "textStyle": "md",
        "px": "4.5"
      },
      "content": {
        "p": "4.5"
      }
    }
  },
  {
    "size": "sm",
    "variant": "line",
    "css": {
      "trigger": {
        "fontSize": "sm",
        "h": "9",
        "minW": "9",
        "px": "2.5"
      },
      "content": {
        "pt": "3"
      }
    }
  },
  {
    "size": "md",
    "variant": "line",
    "css": {
      "trigger": {
        "fontSize": "md",
        "h": "10",
        "minW": "10",
        "px": "3"
      },
      "content": {
        "pt": "4"
      }
    }
  },
  {
    "size": "lg",
    "variant": "line",
    "css": {
      "trigger": {
        "px": "3.5",
        "h": "11",
        "minW": "11",
        "fontSize": "md"
      },
      "content": {
        "pt": "5"
      }
    }
  }
]

const tabsSlotNames = [
  [
    "root",
    "tabs__root"
  ],
  [
    "list",
    "tabs__list"
  ],
  [
    "trigger",
    "tabs__trigger"
  ],
  [
    "content",
    "tabs__content"
  ],
  [
    "indicator",
    "tabs__indicator"
  ]
]
const tabsSlotFns = /* @__PURE__ */ tabsSlotNames.map(([slotName, slotKey]) => [slotName, createRecipe(slotKey, tabsDefaultVariants, getSlotCompoundVariant(tabsCompoundVariants, slotName))])

const tabsFn = (props = {}) => {
  return Object.fromEntries(tabsSlotFns.map(([slotName, slotFn]) => [slotName, slotFn(props)]))
}

const tabsVariantKeys = [
  "variant",
  "size"
]

export const tabs = /* @__PURE__ */ Object.assign(tabsFn, {
  __recipe__: false,
  __name__: 'tabs',
  raw: (props) => props,
  variantKeys: tabsVariantKeys,
  variantMap: {
  "variant": [
    "line",
    "outline"
  ],
  "size": [
    "sm",
    "md",
    "lg"
  ]
},
  splitVariantProps(props) {
    return splitProps(props, tabsVariantKeys)
  },
})