
def update_fields(target, source, exclude=None):
    exclude = exclude or set()

    for field, value in source.model_dump(exclude=exclude).items():
        setattr(target, field, value)