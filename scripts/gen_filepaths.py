import os

transit_plate_images_path = "./data/plates"

brazillian_plate_types = [
    'A',  # [pt-br] Aviso  
    'R',  # [pt-br] Regulamentação (Restrição ou Obrigação)
    'SAU' # [pt-br] Informativo
]

plate_filepaths = [
    os.path.join(transit_plate_images_path, fname) 
    for fname in os.listdir(transit_plate_images_path)
        if fname.split("-")[0] in brazillian_plate_types
]

def load_filepaths():
    for filepath in plate_filepaths:
        print(filepath)

if __name__ == '__main__':
    load_filepaths()